/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:24
 */

package websocket

import (
	"fmt"
	"imserver/lib/cache"
	"imserver/models"
	"strings"
	"sync"
	"time"
)

// 连接管理
type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // teamId+uid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Login       chan *login        // 用户登录处理
	UnLogin     chan *login        // 用户登录处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}

// 获取用户key
func GetUserKey(teamId, userId int) (key string) {
	key = fmt.Sprintf("%d_%d", teamId, userId)

	return
}

/**************************  manager  ***************************************/

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	// 连接存在，在添加
	_, ok = manager.Clients[client]

	return
}

// GetClients
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value

		return true
	})

	return
}

// 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {

	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	for key, value := range manager.Clients {
		result := f(key, value)
		if !result {
			return
		}
	}
}

// GetClientsLen
func (manager *ClientManager) GetClientsLen() (clientsLen int) {

	clientsLen = len(manager.Clients)
	return
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	// if _, ok := manager.Clients[client]; ok {
	delete(manager.Clients, client)
	// }
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(teamId, userId int) (client *Client) {

	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()

	userKey := GetUserKey(teamId, userId)
	if value, ok := manager.Users[userKey]; ok {
		client = value
	}

	return
}

// GetClientsLen
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)

	return
}

// 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	manager.Users[key] = client
}

// 删除用户
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	key := GetUserKey(client.TeamId, client.UserId)
	if value, ok := manager.Users[key]; ok {
		// 判断是否为相同的用户
		if value.Addr != client.Addr {
			return
		}
		delete(manager.Users, key)
		result = true
	}

	return
}

// 获取用户的key
func (manager *ClientManager) GetUserKeys() (userKeys []string) {

	userKeys = make([]string, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for key := range manager.Users {
		userKeys = append(userKeys, key)
	}

	return
}

// 获取用户的key
func (manager *ClientManager) GetUserList(teamId int) (userList []int) {
	userList = make([]int, 0)

	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()

	for k, v := range manager.Users {
		if strings.HasPrefix(k, fmt.Sprintf("%d_", teamId)) {
			userList = append(userList, v.UserId)
		}
	}

	fmt.Println("GetUserList len:", len(manager.Users))

	return
}

// 获取用户的key
func (manager *ClientManager) GetUserClients(teamId int) (clients []*Client) {

	clients = make([]*Client, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()

	for k, v := range manager.Users {
		if strings.HasPrefix(k, fmt.Sprintf("%d_", teamId)) {
			clients = append(clients, v)
		}
	}

	return
}

// 向全部成员(除了自己)发送数据
func (manager *ClientManager) sendTeamIdAll(message []byte, teamId int) {
	clients := manager.GetUserClients(teamId)
	for _, conn := range clients {
		conn.SendMsg(message)
	}

	// 获取没有进入任何群的客户端
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

	// client.Send <- []byte("连接成功")
}

// 用户断开连接
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)

	// 删除用户连接
	if ok := manager.DelUsers(client); !ok {
		// 不是当前连接的客户端
		return
	}

	// 清除redis登录数据
	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err == nil {
		userOnline.LogOut()
		_ = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	}

	// 关闭 chan
	// close(client.Send)

	fmt.Println("EventUnregister 用户断开连接", client.Addr, client.TeamId, client.UserId)

	if client.UserId != 0 {
		// orderId := helper.GetOrderIdTime()
		_, _ = SendUserMessageAll(client.TeamId, client.UserId, 1, models.MessageCmdExit, "用户已经离开~")
	}
}

// 用户离开群
func (manager *ClientManager) EventUnLogin(login *login) {
	client := login.Client
	// 连接存在，在添加
	if manager.InClient(client) {
		userKey := login.GetKey()
		delete(manager.Users, userKey)
	}

	fmt.Println("EventUnLogin 用户离开群", client.Addr, login.TeamId, login.UserId)

	// orderId := helper.GetOrderIdTime()
	_, _ = SendUserMessageAll(login.TeamId, login.UserId, 1, models.MessageCmdUnLogin, "我离开了~")
}

// 用户进群
func (manager *ClientManager) EventLogin(login *login) {

	client := login.Client
	// 连接存在，在添加
	if manager.InClient(client) {
		userKey := login.GetKey()
		manager.AddUsers(userKey, login.Client)
	}

	fmt.Println("EventLogin 用户进入群", client.Addr, login.TeamId, login.UserId)

	// orderId := helper.GetOrderIdTime()
	_, _ = SendUserMessageAll(login.TeamId, login.UserId, 1, models.MessageCmdUnLogin, "哈喽~")
}

// 管道处理程序
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)

		case unlogin := <-manager.UnLogin:
			// 用户退出群
			manager.EventUnLogin(unlogin)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)

		case message := <-manager.Broadcast:
			// 广播事件
			clients := manager.GetClients()
			for conn := range clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
				}
			}
		}
	}
}

/**************************  manager info  ***************************************/
// 获取管理者信息
func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})

	managerInfo["clientsLen"] = clientManager.GetClientsLen()        // 客户端连接数
	managerInfo["usersLen"] = clientManager.GetUsersLen()            // 登录用户数
	managerInfo["chanRegisterLen"] = len(clientManager.Register)     // 未处理连接事件数
	managerInfo["chanLoginLen"] = len(clientManager.Login)           // 未处理登录事件数
	managerInfo["chanUnregisterLen"] = len(clientManager.Unregister) // 未处理退出登录事件数
	managerInfo["chanBroadcastLen"] = len(clientManager.Broadcast)   // 未处理广播事件数

	if isDebug == "true" {
		addrList := make([]string, 0)
		clientManager.ClientsRange(func(client *Client, value bool) (result bool) {
			addrList = append(addrList, client.Addr)

			return true
		})

		users := clientManager.GetUserKeys()

		managerInfo["clients"] = addrList // 客户端列表
		managerInfo["users"] = users      // 登录用户列表
	}

	return
}

// 获取用户所在的连接
func GetUserClient(teamId int, userId int) (client *Client) {
	client = clientManager.GetUserClient(teamId, userId)

	return
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())

	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.UserId, client.LoginTime, client.HeartbeatTime)

			client.Socket.Close()
		}
	}
}

// 获取全部用户
func GetUserList(teamId int) (userList []int) {
	fmt.Println("获取全部用户", teamId)

	userList = clientManager.GetUserList(teamId)

	return
}

// 全员广播
func AllSendMessages(teamId, userId int, data string) {
	// fmt.Println("全员广播", appId, userId, data) ------

	// ignoreClient := clientManager.GetUserClient(teamId, userId)
	clientManager.sendTeamIdAll([]byte(data), teamId)
}
