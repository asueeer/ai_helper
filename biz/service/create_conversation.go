package service

import (
	"ai_helper/biz/common"
	domainService "ai_helper/biz/domain/service"
	"ai_helper/biz/model"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"log"
)

type CreateConversationService struct {
}

func (ss *CreateConversationService) Execute(ctx context.Context, req *model.CreateConversationRequest) (resp *model.CreateConversationResponse, err error) {
	if req.Type == common.HelperConversationType {
		return ss.CreateHelperConv(ctx)
	}
	return nil, common.NewBizErr(common.BizErrCode, "不支持的会话类型", err)
}

func (ss *CreateConversationService) CreateHelperConv(ctx context.Context) (resp *model.CreateConversationResponse, err error) {
	var convDomainService domainService.ConversationService
	var convLoader domainService.ConversationLoader
	user := common.GetUser(ctx)
	{
		// 如果之前已经创建过会话，则直接返回创建的会话id
		convEntity, err := convLoader.GetHelperConvEntityByUserID(ctx, user.UserID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, common.NewBizErr(common.BizErrCode, err.Error(), err)
		}
		if convEntity != nil {

		}
		if convEntity != nil {
			err = convEntity.ReOpen(ctx)
			if err != nil {
				return nil, common.NewBizErr(common.BizErrCode, err.Error(), err)
			}
			return &model.CreateConversationResponse{
				Meta: common.MetaOk,
				Data: model.CreateConversationData{
					ConvID: cast.ToString(convEntity.ConvID),
				},
			}, nil
		}
	}

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
	resp.Data.IsNew = true
	return resp, nil
}
