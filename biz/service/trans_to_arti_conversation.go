package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type TransToArtiConversationService struct {
}

func (s TransToArtiConversationService) Execute(ctx *gin.Context, req *model.TransToArtiConvRequest) (*model.TransToArtiConvResponse, error) {
	if req.ConvID == "" {
		return nil, common.NewBizErr(common.BizErrCode, "conv_id is empty", nil)
	}
	// todo 权限校验
	convRepo := repo.NewConversationRepo()
	// fixme 理论上应该做一个校验, 看有没有被更改
	err := convRepo.UpdateConvStatus(ctx, repo.UpdateConvStatusRequest{
		ConvID:    cast.ToInt64(req.ConvID),
		Status:    common.HelperConvStatusWaiting,
		PreStatus: common.HelperConvStatusRoboting,
	})
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "trans to arti_conv fail", err)
	}

	return &model.TransToArtiConvResponse{
		Meta: common.MetaOk,
		Data: model.CreateConversationData{
			// 有坑 fixme 这个结构体是给长连接发消息用的，并不是每次都需要这样装载信息
			ConvID: cast.ToString(req.ConvID),
			IsNew:  true,
		},
	}, nil
}
