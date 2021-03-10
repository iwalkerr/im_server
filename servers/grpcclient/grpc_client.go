package grpcclient

import (
	"context"
	"errors"
	"fmt"
	"imserver/common"
	"imserver/models"
	"imserver/protobuf"
	"time"

	"google.golang.org/grpc"
)

func SendMsgAll(server *models.Server, seq, teamId, userId int, cmd string, message string) (sendMsgId string, err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgAllReq{
		Seq:    uint32(seq),
		TeamId: uint32(teamId),
		Cms:    cmd,
		UserId: uint32(userId),
		Msg:    message,
	}

	rsp, err := c.SendMsgAll(ctx, &req)
	if err != nil {
		fmt.Println("给全体用户发送消息", err)

		return
	}

	if rsp.GetRetCode() != common.OK {
		fmt.Println("给全体用户发送消息", rsp.String())
		err = errors.New(fmt.Sprintf("发送消息失败 code:%d", rsp.GetRetCode()))
		return
	}

	fmt.Println("给全体用户发送消息 成功")
	return
}
