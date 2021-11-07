package handler

import (
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

// SendMessageCallBack [Post] /im/send_message
func SendMessageCallBack(c *gin.Context) {
	// 发送消息回调
	req := &model.SendMessageRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.SendMessageService
	var err error
	err = ss.ExecuteCallback(c, req)
	if err != nil {
		log.Printf("SendMessageService report, ss.ExecuteCallback: %+v", err)
	}
}
