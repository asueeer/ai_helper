package repo

import (
	"ai_helper/biz/dal/db"
	"ai_helper/biz/dal/db/po"
	"context"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(ctx context.Context) *UserRepo {
	return &UserRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo *UserRepo) Do(ctx context.Context, f func(tx *gorm.DB) error) error {
	return repo.db.Transaction(f)
}

func (repo *UserRepo) GetUserPoByOpenID(ctx context.Context, wxOpenID string) (*po.User, error) {
	var userPo po.User
	sql := repo.db.Model(po.User{})
	sql = sql.Joins("left join user_account on user_center.user_id = user_account.user_id")
	sql = sql.Where("user_account.wx_open_id = ?", wxOpenID).Find(&userPo)
	err := sql.Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 如果是未找到的错误, 则返回(nil, nil)
	}
	if err != nil {
		return nil, err
	}
	return &userPo, nil
}

func (repo *UserRepo) GetUserPoByUserID(ctx context.Context, userID int64) (*po.User, error) {
	var userPo po.User
	sql := repo.db.Model(po.User{})
	sql = sql.Where("user_id = ?", userID).Find(&userPo)
	err := sql.Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 如果是未找到的错误, 则返回(nil, nil)
	}
	if err != nil {
		return nil, err
	}
	return &userPo, nil
}

func (repo *UserRepo) GetUserPos(ctx context.Context, userIDs []int64) ([]*po.User, error) {
	pos := make([]*po.User, 0)
	sql := repo.db.Model(po.User{})
	sql = sql.Where("user_id in (?)", userIDs)
	err := sql.Find(&pos).Error
	if err != nil {
		return nil, err
	}
	return pos, nil
}
