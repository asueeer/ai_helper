package handler

import (
	"github.com/gin-gonic/gin"

	"nearby/biz/common"
	"nearby/biz/model"
	"nearby/biz/service"
)

// SendMessage [Post] /im/send_message
func SendMessage(c *gin.Context) {
	// 发送消息
	req := &model.SendMessageRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.SendMessageService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
