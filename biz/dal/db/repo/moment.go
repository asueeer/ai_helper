package repo

import (
	"ai_helper/biz/dal/db"
	"ai_helper/biz/dal/db/po"
	"context"
	"github.com/jinzhu/gorm"
)

type MomentRepo struct {
	db *gorm.DB
}

func NewMomentRepo(ctx context.Context) *MomentRepo {
	return &MomentRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo MomentRepo) CreateMoment(ctx context.Context, momentPo *po.Moment) (*po.Moment, error) {
	sql := repo.db.Model(po.Moment{})
	err := sql.Omit("id").Create(&momentPo).Error
	if err != nil {
		return nil, err
	}
	return momentPo, nil
}

type GetMomentRequest struct {
	MomentID int64 `json:"moment_id"`
}

func (repo MomentRepo) GetMoment(ctx context.Context, req GetMomentRequest) (*po.Moment, error) {
	var momentPo po.Moment
	sql := repo.db.Model(po.Moment{})
	sql = sql.Where("moment_id = ?", req.MomentID)
	err := sql.First(&momentPo).Error
	if err != nil {
		return nil, err
	}

	return &momentPo, nil
}

type GetMomentsRequest struct {
	Limit          int64  `json:"limit"`
	Offset         int64  `json:"offset"`
	SortByCreateAt string `json:"create_at_desc"`
}

func (repo MomentRepo) GetMoments(ctx context.Context, req GetMomentsRequest) (pos []*po.Moment, total int64, err error) {
	pos = make([]*po.Moment, 0)

	sql := repo.db.Model(po.Moment{})
	sql = sql.Limit(req.Limit)
	sql = sql.Offset(req.Offset)
	if req.SortByCreateAt == Desc {
		sql = sql.Order("moment.create_at desc")
	}
	sql = sql.Order("id desc")
	err = sql.Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	err = sql.Count(&total).Error
	if err != nil {
		return nil, 0, err // 是否要返回错误, 其实还可以再细究一下
	}
	return pos, total, nil
}

func (repo MomentRepo) DeleteMoment(ctx context.Context, momentID int64) error {
	sql := repo.db.Model(po.Moment{})

	sql = sql.Delete(&po.Moment{
		MomentID: momentID,
	}, "moment_id = ?", momentID)

	err := sql.Error
	if err != nil {
		return err
	}

	return nil
}

func (repo MomentRepo) UpdateViewCnt(ctx context.Context, momentPo po.Moment) error {
	sql := repo.db.Model(po.Moment{})
	sql = sql.Where("moment_id = ?", momentPo.MomentID).UpdateColumn("view_count", momentPo.ViewCount)
	return sql.Error
}

func (repo MomentRepo) UpdateCommentCnt(ctx context.Context, momentID int64, bias *int) error {
	sql := repo.db.Model(po.Moment{})
	sql = sql.Where("moment_id = ?", momentID)
	sql = sql.UpdateColumn("comment_count", gorm.Expr("comment_count + ?", bias))
	return sql.Error
}
