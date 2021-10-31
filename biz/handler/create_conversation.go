package handler

import (
	"github.com/gin-gonic/gin"
	"nearby/biz/common"
	"nearby/biz/model"
	"nearby/biz/service"
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
