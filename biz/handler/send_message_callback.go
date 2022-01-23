package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/domain/entity"
	"ai_helper/biz/handler/ws_handler"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

// SendMessageCallBack [Post] /im/send_message
func SendMessageCallBack(c *gin.Context) {
	// 发送消息回调
	req := &model.SendMessageRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	{
		// 更新未读消息数
		var ss service.SendMessageService
		var err error
		err = ss.ExecuteCallback(c, req)
		if err != nil {
			log.Printf("SendMessageService report, ss.ExecuteCallback: %+v", err)
		}
	}

	{
		// 如果用户在线, 就给该长连接发一条消息
		obj := c.Value("msg")
		msg := obj.(*entity.MessageFrom)
		if msg == nil {
			log.Printf("msg is nil")
			return
		}
		ws_handler.TheHub.BatchSendMsgs(c, msg.ReceiverID, model.WsMessageResponse{
			Type: common.WsNewMsg,
			Msg:  msg.ToVo(),
		})
	}
}
