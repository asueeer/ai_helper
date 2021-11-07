package aggregate

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"
	"context"
	"log"
)

// 发送者角度的消息
type MessageAggregate struct {
	MessageID   int64                `json:"id"`
	MessageFrom *entity.MessageFrom  `json:"message_from"`
	MessageTos  []*entity.MessageTo  `json:"message_tos"` // 收件箱的多份信件
	Conv        *entity.Conversation `json:"conv"`
}

func (agg *MessageAggregate) SetMessageID(msgID int64) {
	log.Printf("msgID: %+v", msgID)
	agg.MessageID = msgID
	agg.MessageFrom.MessageID = msgID
}

func (agg *MessageAggregate) SyncReceiverBox(ctx context.Context) error {
	/*
		fixme:
			现在只是单聊, 等做群聊的时候, 收件箱要持久化的数据就多了, 要做用消息队列做异步
			从用户-会话关系表里查找到与该会话有关的用户,
			给这些用户的收件箱发送消息, 持久化起来
	*/
	var err error
	_, err = agg.FindConvEntity(ctx)
	if err != nil {
		return err
	}
	msgTos := agg.ConstructMessageTos(ctx)
	for i := range msgTos {
		err = msgTos[i].Persist(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (agg *MessageAggregate) FindConvEntity(ctx context.Context) (*entity.Conversation, error) {
	convEntity, err := entity.NewConversationEntityByID(ctx, agg.MessageFrom.ConvID)
	if err != nil {
		return nil, err
	}
	agg.Conv = convEntity
	return convEntity, nil
}

func (agg *MessageAggregate) ConstructMessageTos(ctx context.Context) []*entity.MessageTo {
	conv := agg.Conv
	switch conv.Type {
	case common.HelperConversationType:
		return agg.ConstructHelperMsgTos(ctx)
	}
	return nil
}

func (agg *MessageAggregate) ConstructHelperMsgTos(ctx context.Context) []*entity.MessageTo {
	msgFrom := agg.MessageFrom
	conv := agg.Conv
	agg.MessageTos = make([]*entity.MessageTo, 2)
	agg.MessageTos[0] = &entity.MessageTo{
		MessageID: msgFrom.MessageID,
		ConvID:    msgFrom.ConvID,
		OwnerID:   conv.Creator,
		Timestamp: msgFrom.Timestamp,
		HasRead:   entity.MsgHasRead,
	}
	agg.MessageTos[1] = &entity.MessageTo{
		MessageID: msgFrom.MessageID,
		ConvID:    msgFrom.ConvID,
		OwnerID:   common.HelperID,
		Timestamp: msgFrom.Timestamp,
		HasRead:   entity.MsgHasNotRead,
	}
	return agg.MessageTos
}

func NewMessageAggregateByID(ctx context.Context, msgID int64) (*MessageAggregate, error) {
	msgRepo := repo.NewMessageRepo()
	msgAgg := &MessageAggregate{
		MessageID: msgID,
	}
	msgFromPo, err := msgRepo.GetMessageFrom(ctx, repo.GetMessageFromRequest{
		MessageID: msgID,
	})
	if err != nil {
		return nil, err
	}
	msgFromEntity, err := entity.NewMessageFromByPo(msgFromPo)
	if err != nil {
		return nil, err
	}
	msgAgg.MessageFrom = msgFromEntity
	return msgAgg, nil
}
