package grpcserver

import (
	"context"
	"fmt"
	"imserver/common"
	"imserver/models"
	"imserver/protobuf"
	"imserver/servers/websocket"

	"google.golang.org/protobuf/proto"
)

type server struct {
}

func setErr(rsp proto.Message, code uint32, message string) {

	message = common.GetErrorMessage(code, message)
	switch v := rsp.(type) {
	// case *protobuf.QueryUsersOnlineRsp:
	// 	v.RetCode = code
	// 	v.ErrMsg = message
	// case *protobuf.SendMsgRsp:
	// 	v.RetCode = code
	// 	v.ErrMsg = message
	case *protobuf.SendMsgAllRsp:
		v.RetCode = code
		v.ErrMsg = message
	// case *protobuf.GetUserListRsp:
	// 	v.RetCode = code
	// 	v.ErrMsg = message
	default:

	}

}

// 给本机全体用户发消息
func (s *server) SendMsgAll(c context.Context, req *protobuf.SendMsgAllReq) (rsp *protobuf.SendMsgAllRsp, err error) {

	fmt.Println("grpc_request 给本机全体用户发消息", req.String())

	rsp = &protobuf.SendMsgAllRsp{}

	data := models.GetMsgData(int(req.GetSeq()), int(req.GetType()), req.GetCms(), req.GetMsg())

	websocket.AllSendMessages(int(req.GetTeamId()), int(req.GetUserId()), data)

	setErr(rsp, common.OK, "")

	fmt.Println("grpc_response 给本机全体用户发消息:", rsp.String())

	return
}
