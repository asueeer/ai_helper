package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/domain/entity"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"ai_helper/biz/model/vo"
	"context"
	"github.com/spf13/cast"
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
		TimestampFrom: cast.ToInt64(req.Cursor) / 1000,
	})
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, err.Error(), err)
	}
	// 加载最新一条会话
	err = convsLoader.LoadLastMsgs(ctx, convEntities)
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "装载会话里的消息时出错", err)
	}

	// 更新接下来拉取会话时，需要给后端传的cursor
	cursor := convsLoader.Cursor(ctx, convEntities)
	resp := &model.LoadConversationsResponse{
		Meta: common.MetaOk,
		Data: model.LoadConversationData{
			Conversations: ss.ToConvVos(ctx, convEntities),
			NewCursor:     cast.ToString(cursor),
			HasMore:       true,
			Total:         total,
		},
	}

	if len(convEntities) < cast.ToInt(req.Limit) || cast.ToString(cursor) == req.Cursor {
		resp.Data.HasMore = false
		resp.Data.NewCursor = "0"
	}
	return resp, nil
}

func (ss *LoadConversationsService) ToConvVos(ctx context.Context, entities []*entity.Conversation) []*vo.Conversation {
	vos := make([]*vo.Conversation, len(entities))
	for i := range vos {
		vos[i] = &vo.Conversation{
			ConvID:    cast.ToString(entities[i].ConvID),
			Type:      entities[i].Type,
			UnRead:    0,
			Timestamp: entities[i].Timestamp.Unix() * 1000,
		}

		// 加载会话成员信息
		vos[i].Participants = make([]vo.Participant, len(entities[i].Participants))
		for j := range vos[i].Participants {
			vos[i].Participants[j] = vo.Participant{
				UserID:  cast.ToString(entities[i].Participants[j]),
				HeadURL: "",
			}
		}

		if entities[i].LastMsg != nil {
			vos[i].LastMsg = entities[i].LastMsg.ToVo()
		}
	}
	return vos
}
