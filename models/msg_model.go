/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-01
* Time: 10:40
 */

package models

import "imserver/common"

// import "imserver/common"

const (
	MessageCmdMsg     = "msg"
	MessageCmdEnter   = "enter"   // 连接
	MessageCmdExit    = "exit"    // 断开
	MessageCmdLogin   = "login"   // 进群
	MessageCmdUnLogin = "unlogin" // 离群
)

// 消息的定义
type Message struct {
	Type  int    `json:"type"` // 消息类型 1 text/2 img/3 video/4 voice
	Msg   string `json:"msg"`  // 消息内容
	From  int    `json:"from"` // 发送者
	MsgId int    `json:"msgId"`
}

func NewTestMsg(msgType, msgId int, Msg string) (message *Message) {
	message = &Message{
		Type: msgType,
		// From:  from,
		Msg:   Msg,
		MsgId: msgId,
	}

	return
}

func getTextMsgData(msgId, msgType int, cmd, message string) string {
	textMsg := NewTestMsg(msgType, msgId, message)
	head := NewResponseHead(cmd, common.OK, "Ok", textMsg)

	return head.String()
}

// 文本消息
func GetMsgData(msgId, msgType int, cmd string, message string) string {

	return getTextMsgData(msgId, msgType, cmd, message)
}

// 用户进入消息
func GetTextMsgDataEnter(uId, msgId int, message string) string {

	return getTextMsgData(msgId, 1, "enter", message)
}

// 用户退出消息
func GetTextMsgDataExit(uId, msgId int, message string) string {

	return getTextMsgData(msgId, 1, "exit", message)
}
