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
	"log"
)

// AcceptConversation [Post] /im/accept_conversation
func AcceptConversation(c *gin.Context) {
	// 接收会话
	req := &model.AcceptConversationRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.AcceptConversationService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	{
		// 给长连接发送消息
		convAgg, err := aggregate.GetConvAggByID(c, cast.ToInt64(req.ConvID))
		if convAgg == nil || err != nil {
			log.Printf("GetConvAggByID fail, err: %+v", err)
			c.JSON(200, resp)
			return
		}
		log.Printf("aggregate.GetConvAggByID is %+v", convAgg.Conv.Creator)
		wsMsg := convAgg.GetNotifyVisitor(c)
		ws_handler.TheHub.BatchSendMsgs(c, cast.ToInt64(convAgg.Conv.Creator), wsMsg)
	}
	c.JSON(200, resp)
}
