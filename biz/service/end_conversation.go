package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
)

type EndConversationService struct {
}

func (EndConversationService) Execute(ctx context.Context, req *model.EndConversationRequest) (*model.EndConversationResponse, error) {
	if req.ConvID == "" {
		return nil, common.NewBizErr(common.BizErrCode, "conv_id is empty", nil)
	}
	// todo 权限校验
	convRepo := repo.NewConversationRepo()
	err := convRepo.UpdateConvStatus(ctx, repo.UpdateConvStatusRequest{
		ConvID:    cast.ToInt64(req.ConvID),
		Status:    common.HelperConvStatusRoboting,
		PreStatus: common.HelperConvStatusChatting,
		Acceptor:  0,
	})
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "end conv fail", err)
	}
	return &model.EndConversationResponse{
		Meta: common.MetaOk,
	}, nil
}
