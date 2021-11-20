package service

import (
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"
	"ai_helper/biz/util"
	"context"
	"time"
)

type ConversationsLoader struct {
}

type LoadConversationsRequest struct {
	UserID    int64 `json:"user_id"`
	Limit     int64 `json:"limit"`
	SeqIDFrom int64 `json:"seq_id_from"`
	SeqIDTo   int64 `json:"seq_id_to"`
}

// LoadConversations 加载会话列表
func (ss *ConversationsLoader) LoadConversations(ctx context.Context, req LoadConversationsRequest) (entities []*entity.Conversation, total int64, err error) {
	convRepo := repo.NewConversationRepo()
	// 1. 根据条件，找到要加载的convID列表
	convRelPos, total, err := convRepo.GetUserConvRelPos(ctx, repo.GetUserConvRelPosRequest{
		UserID:    req.UserID,
		Limit:     req.Limit,
		SeqIDFrom: req.SeqIDFrom,
		SeqIDTo:   req.SeqIDTo,
	})
	if err != nil {
		return nil, 0, err
	}
	convIDs := ss.ToConvIDs(convRelPos)
	// 2. 根据convID列表找到相关convPo
	convPos, err := convRepo.GetConvPos(ctx, repo.GetConvPosRequest{
		ConvIDs: convIDs,
	})
	convPoMap := make(map[int64]*po.Conversation)
	for i := range convPos {
		convPoMap[convPos[i].ConvID] = convPos[i]
	}

	// 3. 根据convPo组装成convEntity列表
	convEntites := make([]*entity.Conversation, len(convPos))
	for i := range convRelPos {
		convEntites[i] = entity.NewConversationEntityByPo(ctx, convPoMap[convRelPos[i].ConvID])
		convEntites[i].Participants = convRelPos[i].Participants
		convEntites[i].UnreadCnt = convRelPos[i].UnreadCnt
	}
	return convEntites, total, nil
}

func (ss *ConversationsLoader) ToConvIDs(pos []*po.UserConvRel) (convIDs []int64) {
	convIDs = make([]int64, len(pos))
	for i := range convIDs {
		convIDs[i] = pos[i].ConvID
	}
	return convIDs
}

type MsgReq struct {
	ConvID    int64     `json:"conv_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (ss *ConversationsLoader) LoadLastMsgs(ctx context.Context, entities []*entity.Conversation) error {
	msgIDs := make([]int64, len(entities))
	for i := range entities {
		msgIDs[i] = entities[i].LastMsgID
	}
	// 1. 拿到所有的msgFromPos
	msgRepo := repo.NewMessageRepo()
	msgFromPos, err := msgRepo.GetMessageFroms(ctx, repo.GetMessageFromsRequest{
		MessageIDs: msgIDs,
	})
	if err != nil {
		return err
	}

	// 2. 根据pos搞一个map
	record := make(map[int64]*po.MessageFrom)
	for i := range msgFromPos {
		record[msgFromPos[i].MessageID] = msgFromPos[i]
	}

	// 3. 装载entities
	for i := range entities {
		msg := record[entities[i].LastMsgID]
		if msg == nil {
			continue
		}
		entities[i].LastMsg, err = entity.NewMessageFromByPo(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ss *ConversationsLoader) Cursor(ctx context.Context, entities []*entity.Conversation) (cursor int64) {
	cursor = -1
	for i := range entities {
		cursor = util.MaxInt64(entities[i].Timestamp.Unix()*1000, cursor)
	}
	return cursor
}

func (ss *ConversationsLoader) LoadUnreadCnt(ctx context.Context, entities []*entity.Conversation) error {
	return nil
}
