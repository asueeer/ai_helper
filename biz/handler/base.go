package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/handler/ws_handler"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
)

// 给客服发长连接消息
func SendMsg2Helper(ctx context.Context, type_ int, msg interface{}) {
	receiverID := common.HelperID
	resp := model.WsMessageResponse{
		Type: type_,
		Msg:  msg,
	}
	ws_handler.TheHub.BatchSendMsgs(ctx, cast.ToInt64(receiverID), resp)
	给NLP机器人发消息(resp)
}
