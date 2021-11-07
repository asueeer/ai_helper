package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
)

// LoadConversations [Post] /im/load_conversations
func LoadConversations(c *gin.Context) {
	// 加载会话列表
	req := &model.LoadConversationsRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.LoadConversationsService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
