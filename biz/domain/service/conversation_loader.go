package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/aggregate"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"ai_helper/biz/domain/entity"
)

type ConversationLoader struct {
}

type GetConversationRequest struct {
	ConvID    int64
	Limit     int64
	ViewerID  int64
	SeqIDFrom int64
	SeqIDTo   int64
}

type GetConversationResponse struct {
	ConvAgg *entity.Conversation             // 会话聚合根
	MsgAggs []*aggregate.MessageFromReceiver // 接收者角度的消息聚合根列表
}

func (ss *ConversationLoader) GetConversation(ctx context.Context, req GetConversationRequest) (resp *GetConversationResponse, err error) {
	resp = &GetConversationResponse{}
	// 获取会话实体
	convEntity, err := entity.GetConversationEntityByID(ctx, req.ConvID)
	if err != nil {
		return nil, err
	}
	if convEntity == nil {
		return nil, errors.New("未查找到相关会话")
	}
	resp.ConvAgg = convEntity
	// 获取消息列表
	msgResp, err := convEntity.GetMessages(ctx, entity.GetMessagesRequest{
		Limit:     req.Limit,
		SeqIDFrom: req.SeqIDFrom,
		SeqIDTo:   req.SeqIDTo,
		ViewerID:  req.ViewerID,
	})
	if err != nil {
		return nil, err
	}
	// 组合出消息聚合根列表
	resp.MsgAggs, err = ss.ConstructMsgAggs(ctx, msgResp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (ss *ConversationLoader) ConstructMsgAggs(ctx context.Context, resp *entity.GetMessagesResponse) ([]*aggregate.MessageFromReceiver, error) {
	msgTos := resp.MsgTos
	msgFroms := resp.MsgFroms
	msgFromsRecord := ss.ConstructMsgFromsMap(msgFroms)
	if len(msgTos) != len(msgFroms) {
		return nil, errors.New("len(resp.MsgTos) != len(resp.MsgFroms)...")
	}
	ret := make([]*aggregate.MessageFromReceiver, len(resp.MsgTos))
	// 顺序是按照msgTos来的, 不能简单的进行组装, 需要保序
	for i := range msgTos {
		msgFromPo := msgFromsRecord[MessageID(msgTos[i].MessageID)]
		msgFrom, err := entity.NewMessageFromByPo(msgFromPo)
		if err != nil {
			errMsg := fmt.Sprintf("entity.NewMessageFromByPo fail: msgFroms is %+v, err is %+v", msgFroms[i], err)
			return nil, errors.Wrap(err, errMsg)
		}
		msgTo := entity.NewMessageToByPo(msgTos[i])
		ret[i] = &aggregate.MessageFromReceiver{
			MessageID:   msgTos[i].MessageID,
			MessageFrom: *msgFrom,
			MessageTo:   *msgTo,
		}
	}
	return ret, nil
}

type MessageID int64

func (m MessageID) ToInt64() int64 {
	return cast.ToInt64(m)
}

func (ss *ConversationLoader) ConstructMsgFromsMap(msgFromPos []*po.MessageFrom) map[MessageID]*po.MessageFrom {
	ret := make(map[MessageID]*po.MessageFrom)
	for i := range msgFromPos {
		msgFromPo := msgFromPos[i]
		ret[MessageID(msgFromPo.MessageID)] = msgFromPo
	}
	return ret
}

func (ss *ConversationLoader) GetHelperConvEntityByUserID(ctx context.Context, userID int64) (*entity.Conversation, error) {
	convRepo := repo.NewConversationRepo()
	convPo, err := convRepo.GetConvPo(ctx, repo.GetConvPoRequest{
		UserID: userID,
		Type:   common.HelperConversationType,
	})
	if err != nil {
		return nil, err
	}
	if convPo == nil {
		return nil, nil
	}
	return entity.NewConversationEntityByPo(ctx, convPo), nil
}
