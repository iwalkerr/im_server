/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 17:36
 */

package models

import (
	"fmt"
	"time"
)

const (
	heartbeatTimeout = 3 * 60 // 用户心跳超时时间
)

type EditShopAddr struct {
	ShopUserId int    `db:"shop_addr_id" form:"shop_addr_id" binding:"required"`
	Name       string `db:"name" form:"name" binding:"required"`
	Phone      string `db:"phone" form:"phone" binding:"required"`
	Addr       string `db:"addr" form:"addr" binding:"required"`
	Detail     string `db:"detail" form:"detail" binding:"required"`
	IsDefault  int    `db:"is_default" form:"is_default"`
	Uid        int    `db:"user_id" form:"-"`
}

type ShopAddr struct {
	ShopUserId int    `db:"shop_addr_id"`
	Name       string `db:"name" form:"name" binding:"required"`
	Phone      string `db:"phone" form:"phone" binding:"required"`
	Addr       string `db:"addr" form:"addr" binding:"required"`
	Detail     string `db:"detail" form:"detail" binding:"required"`
	IsDefault  int    `db:"is_default" form:"is_default"`
	Uid        int    `db:"user_id" form:"-"`
}

type PwdInfo struct {
	OldPwd string `form:"old_pwd" binding:"required"` // 老密码
	NewPwd string `form:"new_pwd" binding:"required"` // 新密码
}

type UserInfo struct {
	Name     string `db:"name" json:"name" form:"name"`
	RealName string `db:"real_name" json:"real_name" form:"real_name"`
	UserAddr string `db:"user_addr" json:"user_addr" form:"user_addr"`
	Sex      int    `db:"sex" json:"sex" form:"sex"`
	Email    string `db:"email" json:"email" form:"email"`
	Uid      int    `json:"-" form:"-"`
}

// 登陆用户信息
type LoginReq struct {
	Mobile    string `form:"mobile" binding:"required"`
	Password  string `form:"password" binding:"required"`
	SmsCode   string `form:"sms_code" binding:"required"`
	Timestamp int64  `form:"timestamp" binding:"required"`
	Sign      string `form:"sign" binding:"required"`
}

// 注册用户信息
type RegisterReq struct {
	Mobile      string `form:"mobile" binding:"required"`
	Password    string `form:"password" binding:"required"`
	SmsCode     string `form:"sms_code" binding:"required"`
	Name        string `form:"name" binding:"required"`
	Sex         int    `form:"sex" binding:"required"`
	Timestamp   int64  `form:"timestamp" binding:"required"`
	Sign        string `form:"sign" binding:"required"`
	Username    string `form:"-"`
	HeadPicture string `form:"-"` // 设置默认头像 -- 随机
	Salt        string `form:"-"` // 盐值
}

// 用户在线状态
type UserOnline struct {
	AccIp         string `json:"accIp"`         // acc Ip
	AccPort       string `json:"accPort"`       // acc 端口
	AppId         uint32 `json:"appId"`         // appId
	UserId        string `json:"userId"`        // 用户Id
	ClientIp      string `json:"clientIp"`      // 客户端Ip
	ClientPort    string `json:"clientPort"`    // 客户端端口
	LoginTime     uint64 `json:"loginTime"`     // 用户上次登录时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	LogOutTime    uint64 `json:"logOutTime"`    // 用户退出登录的时间
	Qua           string `json:"qua"`           // qua
	DeviceInfo    string `json:"deviceInfo"`    // 设备信息
	IsLogoff      bool   `json:"isLogoff"`      // 是否下线
}

/**********************  数据处理  *********************************/

// 用户登录
func UserLogin(accIp, accPort string, appId uint32, userId string, addr string, loginTime uint64) (userOnline *UserOnline) {

	userOnline = &UserOnline{
		AccIp:         accIp,
		AccPort:       accPort,
		AppId:         appId,
		UserId:        userId,
		ClientIp:      addr,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		IsLogoff:      false,
	}

	return
}

// 用户心跳
func (u *UserOnline) Heartbeat(currentTime uint64) {

	u.HeartbeatTime = currentTime
	u.IsLogoff = false

	return
}

// 用户退出登录
func (u *UserOnline) LogOut() {

	currentTime := uint64(time.Now().Unix())
	u.LogOutTime = currentTime
	u.IsLogoff = true

	return
}

/**********************  数据操作  *********************************/

// 用户是否在线
func (u *UserOnline) IsOnline() (online bool) {
	if u.IsLogoff {

		return
	}

	currentTime := uint64(time.Now().Unix())

	if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
		fmt.Println("用户是否在线 心跳超时", u.AppId, u.UserId, u.HeartbeatTime)

		return
	}

	if u.IsLogoff {
		fmt.Println("用户是否在线 用户已经下线", u.AppId, u.UserId)

		return
	}

	return true
}

// 用户是否在本台机器上
func (u *UserOnline) UserIsLocal(localIp, localPort string) (result bool) {

	if u.AccIp == localIp && u.AccPort == localPort {
		result = true

		return
	}

	return
}
