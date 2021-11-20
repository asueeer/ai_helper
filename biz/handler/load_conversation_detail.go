package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// LoadConversationDetail [Post] /im/load_conversation_detail
func LoadConversationDetail(c *gin.Context) {
	// 加载会话详情
	// 另外一个要做的是让会话未读消息数清零
	req := &model.LoadConversationDetailRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.LoadConversationDetailService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	ctxCopy := c.Copy()
	go func() {
		ss.ClearUnreadCnt(ctxCopy, req.ConvID)
	}()
	c.JSON(200, resp)
}
