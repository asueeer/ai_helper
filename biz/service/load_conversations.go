package service

import (
	"context"
	"github.com/spf13/cast"
	"nearby/biz/common"
	"nearby/biz/domain/entity"
	domainService "nearby/biz/domain/service"
	"nearby/biz/model"
	"nearby/biz/model/vo"
)

type LoadConversationsService struct {
}

func (ss *LoadConversationsService) Execute(ctx context.Context, req *model.LoadConversationsRequest) (*model.LoadConversationsResponse, error) {
	if req.Limit == 0 {
		req.Limit = 40
	}
	convsLoader := domainService.ConversationsLoader{}

	user := common.GetUser(ctx)
	convEntities, total, err := convsLoader.LoadConversations(ctx, domainService.LoadConversationsRequest{
		UserID:        user.UserID,
		Limit:         req.Limit,
		TimestampFrom: cast.ToInt64(req.Cursor),
	})
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, err.Error(), err)
	}
	// 加载最新一条会话
	return &model.LoadConversationsResponse{
		Meta: common.MetaOk,
		Data: model.LoadConversationData{
			Conversations: ss.ToConvVos(ctx, convEntities),
			NewCursor:     "0",
			HasMore:       false,
			Total:         total,
		},
	}, nil
}

func (ss *LoadConversationsService) ToConvVos(ctx context.Context, entities []*entity.Conversation) []*vo.Conversation {
	vos := make([]*vo.Conversation, len(entities))
	for i := range vos {
		vos[i] = &vo.Conversation{
			ConvID:       entities[i].ConvID,
			Type:         entities[i].Type,
			UnRead:       0,
			LastMsg:      vo.Message{},
			Participants: nil,
			ConvIcon:     "",
			Timestamp:    entities[i].Timestamp.Unix(),
		}
	}
	return vos
}
