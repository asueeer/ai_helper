package handler

// CreateConversation [Post] /im/create_conversation
import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func TransToArtiConversation(c *gin.Context) {
	// 创建会话
	req := &model.TransToArtiConvRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.TransToArtiConversationService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	// 给在线客服的长连接里发送消息
	if resp.Data.IsNew {
		SendMsg2Helper(c, common.WsNewConv, resp.Data)
	}

	c.JSON(200, resp)
}
