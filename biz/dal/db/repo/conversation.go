package repo

import (
	"ai_helper/biz/dal/db"
	"ai_helper/biz/dal/db/po"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type ConversationRepo struct {
	db *gorm.DB
}

func NewConversationRepo() *ConversationRepo {
	return &ConversationRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo *ConversationRepo) CreateConversation(ctx context.Context, convPo po.Conversation) (*po.Conversation, error) {
	sql := repo.db.Model(po.Conversation{})
	if err := sql.Omit("id").Create(&convPo).Error; err != nil {
		return nil, err
	}
	return &convPo, nil
}

func (repo *ConversationRepo) CreateUserConvRelPo(ctx context.Context, userConvRelPo po.UserConvRel) (*po.UserConvRel, error) {
	sql := repo.db.Model(po.UserConvRel{})
	if err := sql.Omit("id").Create(&userConvRelPo).Error; err != nil {
		return nil, err
	}
	return &userConvRelPo, nil
}

type GetConvPoRequest struct {
	ConvID int64  `json:"conv_id"`
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
}

func (repo *ConversationRepo) GetConvPo(ctx context.Context, req GetConvPoRequest) (*po.Conversation, error) {
	sql := repo.db.Model(po.Conversation{})
	var conv po.Conversation
	if req.ConvID != 0 {
		sql = sql.Where("conv_id = ?", req.ConvID)
	}
	if req.UserID != 0 {
		sql = sql.Where("creator = ?", req.UserID)
	}
	if req.Type != "" {
		sql = sql.Where("type = ?", req.Type)
	}
	if err := sql.Find(&conv).Error; err != nil {
		return nil, err
	}
	return &conv, nil
}

type GetUserConvRelPosRequest struct {
	ConvID    int64    `json:"conv_id"`
	ConvIDs   []int64  `json:"conv_ids"`
	UserID    int64    `json:"user_id"`
	Limit     int64    `json:"limit"`
	Offset    int64    `json:"offset"`
	SeqIDFrom int64    `json:"seq_id_from"`
	SeqIDTo   int64    `json:"seq_id_to"`
	Status    []string `json:"status"`
	Acceptor  int64    `json:"acceptor"`
}

func (repo *ConversationRepo) GetUserConvRelPos(ctx context.Context, req GetUserConvRelPosRequest) (pos []*po.UserConvRel, total int64, err error) {
	sql := repo.db.Model(po.UserConvRel{})
	pos = make([]*po.UserConvRel, 0)
	if req.ConvID != 0 {
		sql = sql.Where("user_conv_rel.conv_id = ?", req.ConvID)
	}
	sql = sql.Joins("left join conversation on conversation.conv_id = user_conv_rel.conv_id")
	if req.UserID != 0 {
		sql = sql.Where("user_conv_rel.user_id = ?", req.UserID)
	}
	if req.SeqIDTo != 0 {
		sql = sql.Where("conversation.seq_id < ?", req.SeqIDTo)
	}
	if req.SeqIDFrom != 0 {
		sql = sql.Where("conversation.timestamp > ?", req.SeqIDFrom)
	}
	if len(req.ConvIDs) != 0 {
		sql = sql.Where("user_conv_rel.conv_id in (?)", req.ConvIDs)
	}
	if len(req.Status) != 0 {
		sql = sql.Where("conversation.status in (?)", req.Status)
	}
	if req.Acceptor != 0 {
		sql = sql.Where("conversation.acceptor in (0, ?)", req.Acceptor)
	}

	sql = sql.Order("conversation.timestamp desc")
	err = sql.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	sql = sql.Offset(req.Offset)
	sql = sql.Limit(req.Limit)
	err = sql.Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}

type GetConvPosRequest struct {
	ConvIDs []int64 `json:"conv_ids"`
}

func (repo *ConversationRepo) GetConvPos(ctx context.Context, req GetConvPosRequest) ([]*po.Conversation, error) {
	sql := repo.db.Model(po.Conversation{})
	pos := make([]*po.Conversation, 0)
	if len(req.ConvIDs) != 0 {
		sql = sql.Where("conv_id in (?)", req.ConvIDs)
	} else {
		return nil, errors.New("len(req.ConvIDs) == 0")
	}
	err := sql.Find(&pos).Error
	if err != nil {
		return nil, err
	}
	return pos, nil
}

func (repo *ConversationRepo) UpdateLastMsgID(ctx context.Context, convID int64, msgID int64) error {
	sql := repo.db.Model(po.Conversation{})
	sql = sql.Where("conv_id = ?", convID).UpdateColumn("last_msg_id", msgID)
	return sql.Error
}

func (repo *ConversationRepo) UpdateSeqID(ctx context.Context, convID int64, seqID int64) error {
	sql := repo.db.Model(po.Conversation{})
	sql = sql.Where("conv_id = ?", convID).UpdateColumn("seq_id", seqID)
	return sql.Error
}

func (repo *ConversationRepo) UpdateTimestamp(ctx context.Context, convID int64, timestamp time.Time) error {
	sql := repo.db.Model(po.Conversation{})
	sql = sql.Where("conv_id = ?", convID).UpdateColumn("timestamp", timestamp)
	return sql.Error
}

func (repo *ConversationRepo) IncrUnreadCnt(ctx context.Context, convID int64, userID int64) error {
	sql := repo.db.Model(po.UserConvRel{})
	sql = sql.Where("conv_id = ?", convID).Where("user_id = ?", userID)
	sql = sql.UpdateColumn("unread_cnt", gorm.Expr("unread_cnt + 1"))
	return sql.Error
}

func (repo *ConversationRepo) ClearUnreadCnt(ctx context.Context, convID int64, userID int64) error {
	sql := repo.db.Model(po.UserConvRel{})
	sql = sql.Where("conv_id = ?", convID).Where("user_id = ?", userID)
	sql = sql.UpdateColumn("unread_cnt", 0)
	return sql.Error
}

type UpdateConvStatusRequest struct {
	ConvID    int64
	Status    string
	PreStatus string
	Acceptor  int64
}

func (repo *ConversationRepo) UpdateConvStatus(ctx context.Context, req UpdateConvStatusRequest) error {
	sql := repo.db.Model(po.Conversation{})
	sql = sql.Where("conv_id = ?", req.ConvID)
	if req.PreStatus != "" {
		sql = sql.Where("status = ?", req.PreStatus)
	}
	sql = sql.Updates(map[string]interface{}{
		"status":   req.Status,
		"acceptor": req.Acceptor,
	})
	return sql.Error
}
