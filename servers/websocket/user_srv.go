package websocket

import (
	"fmt"
	"imserver/lib/cache"
	"imserver/models"
	"imserver/servers/grpcclient"
	"time"
)

// 发送群消息
func SendUserMessageAll(teamId, userId, msgType int, cmdType, content string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}

	// 存入数据库，返回消息ID，消息时间,用户昵称，头像
	msgId := 1

	for _, server := range servers {
		if IsLocal(server) {
			data := models.GetMsgData(msgId, msgType, cmdType, content)
			AllSendMessages(teamId, userId, data)
		} else {
			_, _ = grpcclient.SendMsgAll(server, msgId, teamId, userId, cmdType, content)
		}
	}

	return
}
