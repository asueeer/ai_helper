package repo

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
	"nearby/biz/dal/db"
	"nearby/biz/dal/db/po"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(ctx context.Context) *CommentRepo {
	return &CommentRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo CommentRepo) CreateComment(ctx context.Context, commentPo po.Comment) (*po.Comment, error) {
	sql := repo.db.Model(po.Comment{})
	err := sql.Omit("id").Create(&commentPo).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &commentPo, nil
}

func (repo CommentRepo) UpdateCommentCnt(ctx context.Context, commentID int64, bias *int) error {
	sql := repo.db.Model(po.Comment{})
	sql = sql.Where("comment_id = ?", commentID)
	if bias != nil {
		sql = sql.UpdateColumn("comment_count = comment_count + ?", bias)
	}
	return sql.Error
}

type GetCommentsRequest struct {
	EntityType string `json:"entity_type"`
	Limit      int64  `json:"limit"`
	Offset     int64  `json:"offset"`
	EntityID   int64  `json:"entity_id"`
}

func (repo CommentRepo) GetComments(ctx context.Context, req GetCommentsRequest) (pos []*po.Comment, total int64, err error) {
	pos = make([]*po.Comment, 0)
	sql := repo.db.Model(po.Comment{})
	if req.EntityType != "" {
		sql = sql.Where("entity_type = ?", req.EntityType)
	}
	if req.EntityID != 0 {
		sql = sql.Where("entity_id = ?", req.EntityID)
	}
	sql = sql.Order("id asc")
	sql = sql.Offset(req.Offset)
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

func (repo CommentRepo) GetComment(ctx context.Context, commentID int64) (*po.Comment, error) {
	var commentPo po.Comment
	sql := repo.db.Model(po.Comment{})
	sql = sql.Where("comment_id = ?", commentID)
	err := sql.First(&commentPo).Error
	if err != nil {
		return nil, err
	}
	return &commentPo, nil
}

func (repo CommentRepo) DeleteMoment(ctx context.Context, commentPo *po.Comment) error {
	sql := repo.db.Model(po.Comment{})
	sql = sql.Where("comment_id = ?", commentPo.CommentID)
	return sql.Delete(commentPo).Error
}
