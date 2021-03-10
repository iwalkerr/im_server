package websocket

import (
	"encoding/json"
	"fmt"
	"imserver/common"
	"imserver/lib/cache"
	"imserver/lib/jwt"
	"imserver/models"
	"time"

	"github.com/go-redis/redis"
)

// ping
func PingController(client *Client, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK
	fmt.Println("webSocket_request ping接口", client.Addr, message)

	data = "pong"
	return
}

// 心跳检测接口
func HeartbeatController(client *Client, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK

	request := &models.HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("心跳接口 解析数据失败", err)
		return
	}

	fmt.Println("webSocket_request 心跳接口", client.TeamId, client.UserId)

	if !client.IsLogin() {
		fmt.Println("心跳接口 用户未登录", client.TeamId, client.UserId)
		code = common.NotLoggedIn
		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = common.NotLoggedIn
			fmt.Println("心跳接口 用户未登录", client.TeamId, client.UserId)
			return
		} else {
			code = common.ServerError
			fmt.Println("心跳接口 GetUserOnlineInfo", client.TeamId, client.UserId, err)
			return
		}
	}

	currentTime := uint64(time.Now().Unix())
	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	if err = cache.SetUserOnlineInfo(client.GetKey(), userOnline); err != nil {
		code = common.ServerError
		fmt.Println("心跳接口 SetUserOnlineInfo", client.TeamId, client.UserId, err)
	}
	return
}

// 用户登录
func LoginController(client *Client, message []byte) (code uint32, msg string, data interface{}) {
	code = common.OK

	request := &models.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("用户登录 解析数据失败", err)
		return
	}

	// 验证token
	uid, _, isValid := jwt.ParseToken(request.ServiceToken)
	if !isValid || uid <= 0 {
		code = common.Unauthorized
		fmt.Println("用户登录 非法的用户")
		return
	}
	request.UserId = uid

	// 验证是否存在该群
	// fmt.Println(request)

	// 是否已经登陆
	if client.IsLogin() {
		fmt.Println("用户登录 用户已经登录", client.UserId)
		code = common.OperationFailure
		return
	}

	currentTime := uint64(time.Now().Unix())
	client.Login(request.TeamId, request.UserId, currentTime)

	// 存储数据
	userOnline := models.UserLogin(serverIp, serverPort, request.TeamId, request.UserId, client.Addr, currentTime)
	if err := cache.SetUserOnlineInfo(client.GetKey(), userOnline); err != nil {
		code = common.ServerError
		fmt.Println("用户登录 SetUserOnlineInfo", err)
		return
	}

	// 用户登录
	login := &login{
		TeamId: request.TeamId,
		UserId: request.UserId,
		Client: client,
	}
	clientManager.Login <- login

	fmt.Println("用户登录 成功", client.Addr, request.UserId)

	return
}
