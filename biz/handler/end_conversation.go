package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/domain/aggregate"
	"ai_helper/biz/handler/ws_handler"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
)

// EndConversation [Post] /im/end_conversation
func EndConversation(c *gin.Context) {
	// 结束会话
	req := &model.EndConversationRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.EndConversationService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	{
		// 给长连接发送消息
		convAgg, err := aggregate.GetConvAggByID(c, cast.ToInt64(req.ConvID))
		if err != nil {
			common.WriteError(c, err)
		}
		wsMsg := convAgg.GetNotifyVisitor(c)
		ws_handler.TheHub.BatchSendMsgs(c, cast.ToInt64(convAgg.Conv.Creator), wsMsg)
	}
	c.JSON(200, resp)
}
