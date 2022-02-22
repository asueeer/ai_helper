package aggregate

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"
	"ai_helper/biz/model"
	"context"
)

type ConversationAggregate struct {
	ConvID   int64                `json:"conv_id"`
	Conv     *entity.Conversation `json:"conversation"`
	ConvRels []*po.UserConvRel    `json:"conv_rels"`
}

type NotifyVisitorMsg struct {
	ConvID int64 `json:"conv_id"`
}

func (agg ConversationAggregate) GetNotifyVisitor(ctx context.Context) model.WsMessageResponse {
	// 给在线客服的长连接里发送消息
	wsMsg := model.WsMessageResponse{
		Type: 0,
		Msg: NotifyVisitorMsg{
			ConvID: agg.ConvID,
		},
	}
	if agg.Conv.Status == common.HelperConvStatusChatting {
		wsMsg.Type = 102
	}
	if agg.Conv.Status == common.HelperConvStatusRoboting {
		wsMsg.Type = 102
	}
	return wsMsg
}

func GetConvAggByID(ctx context.Context, convID int64) (*ConversationAggregate, error) {
	convEntity, err := entity.GetConversationEntityByID(ctx, convID)
	if err != nil {
		return nil, err
	}
	if convEntity == nil {
		return nil, nil
	}
	convRepo := repo.NewConversationRepo()
	convRelPos, _, err := convRepo.GetUserConvRelPos(ctx, repo.GetUserConvRelPosRequest{
		ConvID: convID,
		Limit:  500, // 魔法数，等着要改
	})
	if err != nil {
		return nil, err
	}
	return &ConversationAggregate{
		ConvID:   convID,
		Conv:     convEntity,
		ConvRels: convRelPos,
	}, nil
}
