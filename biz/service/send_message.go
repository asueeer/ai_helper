package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/aggregate"
	"ai_helper/biz/domain/entity"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"ai_helper/biz/util"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"log"
	"time"
)

type SendMessageService struct {
}

// 当前用户是否有权限发送消息？
func (ss *SendMessageService) checkParams(ctx context.Context, req *model.SendMessageRequest) (err error) {
	var convAgg *aggregate.ConversationAggregate
	// 1. 先看有没有这个会话
	convAgg, err = aggregate.GetConvAggByID(ctx, cast.ToInt64(req.ConvID))
	if err != nil {
		return err
	}
	if convAgg == nil {
		return errors.New("不存在的会话id")
	}

	// 2. 再看用户具不具有发送消息的资格
	hasAuth := false
	user := common.GetUser(ctx)
	for i := range convAgg.ConvRels {
		convRel := convAgg.ConvRels[i]
		log.Printf("convRel: %+v", convRel)
		log.Printf("user:  %+v", user)
		if convRel.UserID == user.UserID {
			hasAuth = true
		}
	}
	// 还有一种情况没有考虑, 如果当前用户是客服, 则无需进行这种发送消息的权限校验
	if !hasAuth {
		return errors.New("似乎没有这个会话的权限????")
	}

	return nil
}

func (ss *SendMessageService) Execute(ctx *gin.Context, req *model.SendMessageRequest) (*model.SendMessageResponse, error) {
	if req.TimestampMs == 0 {
		req.TimestampMs = util.Sec2Mirco(time.Now().Unix())
	}
	if err := ss.checkParams(ctx, req); err != nil {
		return nil, common.NewBizErr(common.BizErrCode, err.Error(), err)
	}
	msgFromEntity, err := ss.SendMsg(ctx, req)
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "ops, 消息没发出去...", err)
	}

	ctx.Set("msg", msgFromEntity) // 给后面的callback函数进行调用
	return &model.SendMessageResponse{
		Meta: common.MetaOk,
		Data: model.SendMessageData{
			MessageID: cast.ToString(msgFromEntity.MessageID),
			ConvID:    cast.ToString(msgFromEntity.ConvID),
		},
	}, nil
}

func (ss *SendMessageService) SendMsg(ctx context.Context, req *model.SendMessageRequest) (*entity.MessageFrom, error) {
	// attention 这里直接调用了消息领域服务,
	// 如果后面消息系统变得很重, 可以引入消息队列进行异步、解耦、削峰
	// 很重的意思是指:
	// 1. 由于消息是写扩散的, 发消息是一个很耗时的操作, 如果是群聊的话，前端等不起, 引入消息可以使其异步
	// 2. 除了发消息, 业务上还要通知、发站内信等, 可以进行业务解耦
	// 3. 发消息的人很多, 消息服务器处理不过来, 引入消息队列可以有削峰的效果
	msgJson, err := json.Marshal(req.Content)
	if err != nil {
		return nil, err
	}
	var messageDomainService domainService.MessageService
	sendMsgReq := domainService.SendMessageRequest{
		ConvID:    cast.ToInt64(req.ConvID),
		Role:      req.Role,
		Content:   msgJson,
		Type:      req.Type,
		Status:    req.Status,
		Timestamp: util.Micro2Sec(req.TimestampMs),
		SeqID:     req.TimestampMs,
	}
	msgFromEntity, err := messageDomainService.SendMessage(ctx, sendMsgReq)
	return msgFromEntity, err
}

func (ss *SendMessageService) ExecuteCallback(ctx context.Context, req *model.SendMessageRequest) error {
	convRepo := repo.NewConversationRepo()

	obj := ctx.Value("msg")
	if obj == nil {
		return errors.New("obj is nil")
	}
	msg := obj.(*entity.MessageFrom)
	if msg == nil {
		return errors.New("msg is nil")
	}
	if err := convRepo.UpdateLastMsgID(ctx, cast.ToInt64(req.ConvID), msg.MessageID); err != nil {
		return err
	}
	if err := convRepo.UpdateSeqID(ctx, cast.ToInt64(req.ConvID), msg.SeqID); err != nil {
		return err
	}
	if err := convRepo.UpdateTimestamp(ctx, cast.ToInt64(req.ConvID), msg.CreateAt); err != nil {
		return err
	}
	if err := convRepo.IncrUnreadCnt(ctx, cast.ToInt64(req.ConvID), msg.ReceiverID); err != nil {
		return err
	}
	return nil
}
