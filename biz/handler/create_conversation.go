package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
)

// CreateConversation [Post] /im/create_conversation
func CreateConversation(c *gin.Context) {
	// 创建会话
	req := &model.CreateConversationRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.CreateConversationService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
