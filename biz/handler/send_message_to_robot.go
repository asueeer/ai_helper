package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SendMessageToRobot [Post] /api/im/send_message_to_robot
func SendMessageToRobot(c *gin.Context) {
	// 给机器人发送消息
	req := &model.SendMessageRequest{}
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.SendMessageToRobotService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	c.JSON(200, resp)
}