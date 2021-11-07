package service

import (
	"ai_helper/biz/common"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"context"
	"github.com/spf13/cast"
	"log"
)

type CreateConversationService struct {
}

func (ss *CreateConversationService) Execute(ctx context.Context, req *model.CreateConversationRequest) (resp *model.CreateConversationResponse, err error) {
	return ss.CreateHelperConv(ctx)
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
