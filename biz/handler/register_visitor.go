package handler

import (
	"github.com/gin-gonic/gin"
	"nearby/biz/common"
	"nearby/biz/model"
	"nearby/biz/service"
)

// RegisterVisitor [Post] /get_token
func RegisterVisitor(c *gin.Context) {
	req := &model.RegisterVisitorRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var ss service.RegisterVisitorService
	resp, err := ss.Execute(c, req)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	c.JSON(200, resp)
}
