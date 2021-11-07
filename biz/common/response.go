package common

import (
	"ai_helper/biz/model"

	"github.com/gin-gonic/gin"
)

var MetaOk = model.Meta{
	Code: 0,
	Msg:  "success",
}

// WriteError 通用的biz err
func WriteError(c *gin.Context, err error) {
	c.JSON(
		200,
		model.Response{
			Meta: model.Meta{
				Code: GetErrorCode(err),
				Msg:  GetErrorMsg(err),
			},
			Data: nil,
		},
	)
	c.Abort()
}
