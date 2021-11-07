package handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/model"
	"ai_helper/biz/service"

	"github.com/gin-gonic/gin"
)

// ProfileMe [Get] /profile/me
func ProfileMe(c *gin.Context) {
	req := &model.ProfileMeRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.ProfileMeService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
