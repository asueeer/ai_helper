package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
)

type AcceptConversationService struct {
}

func (AcceptConversationService) Execute(ctx context.Context, req *model.AcceptConversationRequest) (*model.AcceptConversationResponse, error) {
	if req.ConvID == "" {
		return nil, common.NewBizErr(common.BizErrCode, "conv_id is empty", nil)
	}
	// todo 权限校验
	convRepo := repo.NewConversationRepo()
	user := common.GetUser(ctx)
	// fixme 理论上应该做一个校验, 看有没有被更改
	err := convRepo.UpdateConvStatus(ctx, repo.UpdateConvStatusRequest{
		ConvID:    cast.ToInt64(req.ConvID),
		Status:    common.HelperConvStatusChatting,
		PreStatus: common.HelperConvStatusWaiting,
		Acceptor:  cast.ToInt64(user.UserID),
	})
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "accept conv fail", err)
	}

	return &model.AcceptConversationResponse{
		Meta: common.MetaOk,
	}, nil
}
