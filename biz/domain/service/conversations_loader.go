package service

import (
	"context"
	"fmt"
	"nearby/biz/dal/db/po"
	"nearby/biz/dal/db/repo"
	"nearby/biz/domain/entity"
)

type ConversationsLoader struct {
}

type LoadConversationsRequest struct {
	UserID        int64 `json:"user_id"`
	Limit         int64 `json:"limit"`
	TimestampFrom int64 `json:"timestamp_from"`
	TimestampTo   int64 `json:"timestamp_to"`
}

// LoadConversations 加载会话列表
func (ss *ConversationsLoader) LoadConversations(ctx context.Context, req LoadConversationsRequest) (entities []*entity.Conversation, total int64, err error) {
	convRepo := repo.NewConversationRepo()
	// 1. 找到要加载的会话convID列表
	convRelPos, total, err := convRepo.GetUserConvRelPos(ctx, repo.GetUserConvRelPosRequest{
		UserID:        req.UserID,
		Limit:         req.Limit,
		TimestampFrom: req.TimestampFrom,
		TimestampTo:   req.TimestampTo,
	})
	if err != nil {
		return nil, 0, err
	}
	convIDs := ss.ToConvIDs(convRelPos)
	// 2. 根据convID列表加载会话详情
	// todo
	fmt.Println(convIDs)
	return nil, 0, nil
}

func (ss *ConversationsLoader) ToConvIDs(pos []*po.UserConvRel) (convIDs []int64) {
	convIDs = make([]int64, len(pos))
	for i := range convIDs {
		convIDs[i] = pos[i].ConvID
	}
	return convIDs
}
