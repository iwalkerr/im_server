/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-01
* Time: 10:40
 */

package models

import "imserver/common"

const (
	// MessageTypeText  = "text"  // 文字
	// MessageTypeImg   = "img"   // 相片
	// MessageTypeVideo = "video" // 视频
	// MessageTypeVoice = "voice" // 语音

	MessageCmdMsg   = "msg"
	MessageCmdEnter = "enter"
	MessageCmdExit  = "exit"
)

// 消息的定义
type Message struct {
	// Target string `json:"target"` // 目标
	Type int    `json:"type"` // 消息类型 1 text/2 img/3 video/4 voice
	Msg  string `json:"msg"`  // 消息内容
	From int    `json:"from"` // 发送者
}

func NewTestMsg(from, msgType int, Msg string) (message *Message) {

	message = &Message{
		Type: msgType,
		From: from,
		Msg:  Msg,
	}

	return
}

func getTextMsgData(uuId, msgId, msgType int, cmd, message string) string {
	textMsg := NewTestMsg(uuId, msgType, message)
	head := NewResponseHead(msgId, cmd, common.OK, "Ok", textMsg)

	return head.String()
}

// 文本消息
func GetMsgData(uId, msgId, msgType int, cmd, message string) string {

	return getTextMsgData(uId, msgId, msgType, cmd, message)
}

// // 文本消息
// func GetTextMsgData(uuId, msgId, message string) string {

// 	return getTextMsgData("msg", uuId, msgId, message)
// }

// 用户进入消息
func GetTextMsgDataEnter(uId, msgId int, message string) string {

	return getTextMsgData(uId, msgId, 1, "enter", message)
}

// 用户退出消息
func GetTextMsgDataExit(uId, msgId int, message string) string {

	return getTextMsgData(uId, msgId, 1, "exit", message)
}
