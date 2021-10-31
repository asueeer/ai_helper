package service

import (
	"context"
	"github.com/spf13/cast"
	"log"
	"nearby/biz/common"
	domainService "nearby/biz/domain/service"
	"nearby/biz/model"
)

type CreateConversationService struct {
}

func (ss *CreateConversationService) Execute(ctx context.Context, req *model.CreateConversationRequest) (resp *model.CreateConversationResponse, err error) {
	switch req.Type {
	// 不同的会话类型创建逻辑不同
	// 客服会话: 不需要在req中填写ReceiverID, 固定ReceiverID为小助手的用户id
	case common.HelperConversationType:
		return ss.CreateHelperConv(ctx)
	default:
		return nil, common.NewBizErr(common.BizErrCode, "不支持的会话类型", nil)
	}
}

func (ss *CreateConversationService) CreateHelperConv(ctx context.Context) (resp *model.CreateConversationResponse, err error) {
	var convDomainService domainService.ConversationService
	// fixme convDomainService查看是否已有符合条件的会话，如果有，则报错+返回会话id
	entity, err := convDomainService.CreateHelperConversation(ctx)
	if err != nil {
		log.Printf("convDomainService.CreateHelperConversation fail, err:%+v", err)
		return nil, common.NewBizErr(common.BizErrCode, "创建会话失败", err)
	}
	resp = &model.CreateConversationResponse{
		Meta: common.MetaOk,
		Data: model.CreateConversationData{
			ConvID: cast.ToString(entity.ConvID),
		},
	}
	return resp, nil
}
