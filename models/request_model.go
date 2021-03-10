/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:41
 */

package models

/************************  请求数据  **************************/
// 通用请求数据格式
type Request struct {
	// Seq  int         `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// 登录请求数据
type Login struct {
	ServiceToken string `json:"token"` // 验证用户是否登录
	TeamId       int    `json:"teamId,omitempty"`
	UserId       int    `json:"-"`
}

// 心跳请求数据
type HeartBeat struct {
	ServiceToken string `json:"token,omitempty"`
}
