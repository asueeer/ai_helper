package repo

import (
	"ai_helper/biz/dal/db"
	"ai_helper/biz/dal/db/po"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo *MessageRepo) CreateMessageFrom(ctx context.Context, messageFrom po.MessageFrom) (*po.MessageFrom, error) {
	sql := repo.db.Model(&po.MessageFrom{})
	if err := sql.Omit("id").Create(&messageFrom).Error; err != nil {
		return nil, err
	}
	return &messageFrom, nil
}

func (repo *MessageRepo) CreateMessageTo(ctx context.Context, messageTo po.MessageTo) (*po.MessageTo, error) {
	sql := repo.db.Model(&po.MessageTo{})
	if err := sql.Omit("id").Create(&messageTo).Error; err != nil {
		return nil, err
	}
	return &messageTo, nil
}

type GetMessageFromRequest struct {
	MessageID int64 `json:"message_id"`
}

func (repo *MessageRepo) GetMessageFrom(ctx context.Context, req GetMessageFromRequest) (msgFrom *po.MessageFrom, err error) {
	sql := repo.db.Model(&po.MessageFrom{})
	if req.MessageID != 0 {
		sql.Where("message_id = ?", req.MessageID)
	}

	err = sql.Find(msgFrom).Error
	if err != nil {
		return nil, err
	}
	return msgFrom, nil
}

type GetMessageFromsRequest struct {
	ConvID     int64   `json:"conv_id"`
	MessageIDs []int64 `json:"message_ids"`
}

func (repo *MessageRepo) GetMessageFroms(ctx context.Context, req GetMessageFromsRequest) (msgFroms []*po.MessageFrom, err error) {
	sql := repo.db.Model(po.MessageFrom{})
	if len(req.MessageIDs) != 0 {
		sql = sql.Where("message_id in (?)", req.MessageIDs)
	}
	if req.ConvID != 0 {
		sql = sql.Where("conv_id = ?", req.ConvID)
	}
	err = sql.Find(&msgFroms).Error
	if err != nil {
		return nil, errors.Wrap(err, "sql.Find(msgFroms) fail")
	}

	return msgFroms, err
}

type GetMessagesRequest struct {
	ConvID    int64 `json:"conv_id"`
	Limit     int64 `json:"limit"`
	SeqIDFrom int64 `json:"seq_id_from"`
	SeqIDTo   int64 `json:"seq_id_to"`
	OwnerID   int64 `json:"owner_id"`
}

func (repo *MessageRepo) GetMessageTos(ctx context.Context, req GetMessagesRequest) (pos []*po.MessageTo, total int64, err error) {
	sql := repo.db.Model(po.MessageTo{})
	if req.ConvID != 0 {
		sql = sql.Where("conv_id = ?", req.ConvID)
	}
	if req.SeqIDFrom != 0 {
		sql = sql.Where("seq_id > ?", req.SeqIDFrom)
	}
	if req.SeqIDTo != 0 {
		sql = sql.Where("seq_id < ?", req.SeqIDTo)
	}
	if req.OwnerID != 0 {
		sql = sql.Where("owner_id = ?", req.OwnerID)
	}
	sql = sql.Order("seq_id desc")
	sql = sql.Limit(req.Limit)
	err = sql.Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	err = sql.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}
