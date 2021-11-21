package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/handler/ws_handler"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
)

// CreateConversation [Post] /im/create_conversation
func CreateConversation(c *gin.Context) {
	// 创建会话
	req := &model.CreateConversationRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.CreateConversationService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	{
		// 给在线客服的长连接里发送消息
		if resp.Data.IsNew && req.Type == common.HelperConversationType {
			receiverID := common.HelperID
			ws_handler.TheHub.BatchSendMsgs(c, cast.ToInt64(receiverID), model.WsMessageResponse{
				Type: 101,
				Msg:  resp.Data,
			})
		}
	}

	c.JSON(200, resp)
}
