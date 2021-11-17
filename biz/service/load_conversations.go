package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/domain/entity"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"ai_helper/biz/model/vo"
	"ai_helper/biz/util"
	"context"
	"github.com/spf13/cast"
)

type LoadConversationsService struct {
}

func (ss *LoadConversationsService) Execute(ctx context.Context, req *model.LoadConversationsRequest) (*model.LoadConversationsResponse, error) {
	if req.Limit == 0 {
		req.Limit = 40
	}
	if req.Direction == 0 {
		req.Direction = -1
	}
	if req.Cursor == "" {
		req.Cursor = cast.ToString(util.NowUnixMicro())
	}
	convsLoader := domainService.ConversationsLoader{}

	loadConvReq := ss.GetLoadConvReq(ctx, req)
	convEntities, total, err := convsLoader.LoadConversations(ctx, loadConvReq)
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

func (ss *LoadConversationsService) ToConvVos(ctx context.Context, convEntities []*entity.Conversation) []*vo.Conversation {
	vos := make([]*vo.Conversation, len(convEntities))
	for i := range vos {
		vos[i] = convEntities[i].ToVo()

		// 加载会话成员信息
		vos[i].Participants = make([]vo.Participant, len(convEntities[i].Participants))
		for j := range vos[i].Participants {
			vos[i].Participants[j] = vo.Participant{
				UserID:  cast.ToString(convEntities[i].Participants[j]),
				HeadURL: "", // 补充头像信息
			}
		}

		if convEntities[i].LastMsg != nil {
			vos[i].LastMsg = convEntities[i].LastMsg.ToVo()
		}
	}
	return vos
}

func (ss *LoadConversationsService) GetLoadConvReq(ctx context.Context, req *model.LoadConversationsRequest) domainService.LoadConversationsRequest {
	user := common.GetUser(ctx)
	loadConvReq := domainService.LoadConversationsRequest{
		UserID: user.UserID,
		Limit:  req.Limit,
	}
	if req.Direction == -1 {
		loadConvReq.SeqIDTo = cast.ToInt64(req.Cursor)
	}
	if req.Direction == 1 {
		loadConvReq.SeqIDFrom = cast.ToInt64(req.Cursor)
	}
	return loadConvReq
}
