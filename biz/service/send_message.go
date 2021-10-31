package service

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"log"
	"nearby/biz/common"
	"nearby/biz/domain/entity"
	domainService "nearby/biz/domain/service"
	"nearby/biz/model"
)

type SendMessageService struct {
}

func (ss *SendMessageService) checkParams(ctx context.Context, req *model.SendMessageRequest) error {
	// todo 权限校验
	// 是否当前用户有权限发送消息？
	return nil
}

func (ss *SendMessageService) Execute(ctx context.Context, req *model.SendMessageRequest) (*model.SendMessageResponse, error) {
	if err := ss.checkParams(ctx, req); err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "ops, 发送消息出错了", err)
	}
	log.Printf("req: %+v, req.Content: %+v", req, req.Content)
	msgFromEntity, err := ss.SendMsg(ctx, req)
	if err != nil {
		return nil, common.NewBizErr(common.BizErrCode, "ops, 消息没发出去...", err)
	}

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
	msgFromEntity, err := messageDomainService.SendMessage(ctx, domainService.SendMessageRequest{
		ConvID:     cast.ToInt64(req.ConvID),
		Role:       req.Role,
		ReceiverID: cast.ToInt64(req.ReceiverID),
		Content:    msgJson,
		Type:       req.Type,
		Status:     req.Status,
		Timestamp:  req.Timestamp,
	})
	return msgFromEntity, err
}
