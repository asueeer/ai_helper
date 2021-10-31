package service

import (
	"context"
	"nearby/biz/common"
	domainService "nearby/biz/domain/service"
	"nearby/biz/model"
)

type LoadConversationsService struct {
}

func (ss *LoadConversationsService) Execute(ctx context.Context, req *model.LoadConversationsRequest) (*model.LoadConversationsResponse, error) {
	convsLoader := domainService.ConversationsLoader{}

	user := common.GetUser(ctx)
	convsLoader.LoadConversations(ctx, domainService.LoadConversationsRequest{
		UserID:        user.UserID,
		Limit:         req.Limit,
		TimestampFrom: req.Cursor,
	})
	// 加载最新一条会话
	return nil, nil
}
