package handler

import (
	"nearby/biz/common"
	"nearby/biz/model"
	"nearby/biz/service"

	"github.com/gin-gonic/gin"
)

// LoadConversationDetail [Post] /im/load_conversation_detail
func LoadConversationDetail(c *gin.Context) {
	// 加载会话详情
	req := &model.LoadConversationDetailRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.LoadConversationDetailService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
