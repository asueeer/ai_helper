package repo

import (
	"context"
	"nearby/biz/common"
	"nearby/biz/dal/db"
	"nearby/biz/dal/db/po"

	"github.com/jinzhu/gorm"
)

type BizKVRepo struct {
	db *gorm.DB
}

func NewBizKVRepo(ctx context.Context) *BizKVRepo {
	return &BizKVRepo{
		db: db.GetDB().Debug().LogMode(true),
	}
}

func (repo *BizKVRepo) GetBizKV(ctx context.Context, key string) (*po.BizKV, error) {
	var bizKV po.BizKV
	sql := repo.db.Model(po.BizKV{})
	sql = sql.Where("biz_kv.key = ?", key).Find(&bizKV)
	err := sql.Error
	if err != nil {
		return nil, err
	}
	return &bizKV, nil
}

func (repo *BizKVRepo) GetVal(ctx context.Context, key string) string {
	bizKv, err := repo.GetBizKV(ctx, key)
	if err != nil || bizKv == nil {
		return ""
	}
	return bizKv.Val
}

func (repo *BizKVRepo) GetKeyCnt(ctx context.Context, pattern string) (cnt int, err error) {
	sql := repo.db.Model(po.BizKV{})
	sql = sql.Where("key like ?", "%"+pattern+"%")
	err = sql.Count(&cnt).Error
	if err != nil {
		return 0, nil
	}
	return cnt, nil
}

func (repo *BizKVRepo) GetDefaultHeadURLCnt(ctx context.Context) (cnt int, err error) {
	return repo.GetKeyCnt(ctx, common.RandomHeadUrlPrefix)
}

func (repo *BizKVRepo) GetLuckyNameCnt(ctx context.Context) (cnt int, err error) {
	return repo.GetKeyCnt(ctx, common.RandomNicknamePrefix)
}

func (repo *BizKVRepo) GetLuckyNumberCnt(ctx context.Context) (cnt int, err error) {
	return repo.GetKeyCnt(ctx, common.RandomLuckyNumberPrefix)
}
func (repo *BizKVRepo) GetRandomDelimiterCnt(ctx context.Context) (cnt int, err error) {
	return repo.GetKeyCnt(ctx, common.RandomDelimiterPrefix)
}
