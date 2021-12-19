package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/dal/db/po"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"
	"ai_helper/biz/domain/val_obj"
	"context"
	"time"

	"github.com/pkg/errors"
)

// ConversationService 会话领域上下文服务
type ConversationService struct {
}

// CreateHelperConversation 创建客服会话
func (ss *ConversationService) CreateHelperConversation(ctx context.Context) (*entity.Conversation, error) {
	user := common.GetUser(ctx)
	// 1. 创建一条会话实体
	convRepo := repo.NewConversationRepo()
	convEntity := ss.ConstructConversationEntity(ctx,
		ConstructConversationEntityRequest{
			user:     user,
			convType: common.HelperConversationType,
			status:   common.HelperConvStatusRoboting,
		},
	)
	// 将该会话持久化到数据库
	_, err := convRepo.CreateConversation(ctx, convEntity.ToPo())
	if err != nil {
		return nil, errors.Wrap(err, "CreateConversation fail.")
	}

	// 2. 维护两个用户-会话关系: 创建者-会话的关系、小助手和该会话的关系
	relPos := ss.ConstructUserConvRel(ctx, convEntity)
	for i := range relPos {
		// 将用户-会话关系表存储到数据库
		_, err = convRepo.CreateUserConvRelPo(ctx, *relPos[i])
		if err != nil {
			return nil, errors.Wrap(err, "CreateUserConvRelPo fail.")
		}
	}
	return &convEntity, nil
}

type ConstructConversationEntityRequest struct {
	user     *val_obj.UserClaims
	convType string
	status   string
}

func (ss *ConversationService) ConstructUserConvRel(ctx context.Context, convEntity entity.Conversation) []*po.UserConvRel {
	pos := make([]*po.UserConvRel, 0)
	pos = append(pos, &po.UserConvRel{
		RelID:        config.GenerateIDInt64(),
		ConvID:       convEntity.ConvID,
		UserID:       convEntity.Creator,
		Role:         common.ConvRoleCreator,
		Participants: []int64{convEntity.Creator, common.HelperID},
	})
	pos = append(pos, &po.UserConvRel{
		RelID:        config.GenerateIDInt64(),
		ConvID:       convEntity.ConvID,
		UserID:       common.HelperID,
		Role:         common.ConvRoleHelper,
		Participants: []int64{convEntity.Creator, common.HelperID},
	})
	return pos
}

// ConstructConversationEntity 构造会话实体
func (ss *ConversationService) ConstructConversationEntity(ctx context.Context, req ConstructConversationEntityRequest) entity.Conversation {
	user := req.user
	return entity.Conversation{
		ConvID:    config.GenerateIDInt64(),
		Type:      req.convType,
		Creator:   user.UserID,
		Status:    req.status,
		Timestamp: time.Now(),
	}
}
