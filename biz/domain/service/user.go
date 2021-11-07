package service

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/spf13/cast"
)

// UserService 用户领域上下文服务
type UserService struct {
}

// GetUserEntityByOpenID
// 根据open_id获取用户实体
// 如果查询失败, 返回错误;
// 如果没找到, 则返回(nil, nil), 不报错
func (ss *UserService) GetUserEntityByOpenID(ctx context.Context, wxOpenID string) (user *entity.User, err error) {
	if wxOpenID == "" {
		return nil, errors.New("wxOpenID is empty")
	}
	userRepo := repo.NewUserRepo(ctx)
	userPo, err := userRepo.GetUserPoByOpenID(ctx, wxOpenID)
	if err != nil {
		return nil, err
	}
	if userPo == nil {
		return nil, nil
	}
	userEntity := entity.NewUserEntityByPo(userPo)
	return userEntity, nil
}

// 根据userIDd获取用户实体
// 如果查询失败, 返回错误;
// 如果没找到, 则返回(nil, nil), 不报错
func (ss *UserService) GetUserEntityByUserID(ctx context.Context, userID int64) (user *entity.User, err error) {
	userRepo := repo.NewUserRepo(ctx)
	userPo, err := userRepo.GetUserPoByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userPo == nil {
		return nil, nil
	}
	userEntity := entity.NewUserEntityByPo(userPo)
	return userEntity, nil
}

func (ss *UserService) GenerateRandomHeadURL(ctx context.Context) (string, error) {
	kvRepo := repo.NewBizKVRepo(ctx)
	key := common.RandomHeadUrlPrefix + cast.ToString(rand.Intn(config.DefaultHeadURLCnt))
	bizKv, err := kvRepo.GetBizKV(ctx, key)
	if err != nil {
		log.Printf("kvRepo.GetBizKV fail, err: %+v", err)
		return "", err
	}
	if bizKv == nil {
		return "", errors.New("bizKv is nil")
	}
	return bizKv.Val, nil
}

func (ss *UserService) GenerateRandomNickname(ctx context.Context) (string, error) {
	kvRepo := repo.NewBizKVRepo(ctx)
	nameKey := common.RandomNicknamePrefix + cast.ToString(rand.Intn(config.LuckyNameCnt))
	luckName := kvRepo.GetVal(ctx, nameKey)
	numKey := common.RandomLuckyNumberPrefix + cast.ToString(rand.Intn(config.LuckyNumberCnt))
	luckNum := kvRepo.GetVal(ctx, numKey)
	delimiterKey := common.RandomDelimiterPrefix + cast.ToString(rand.Intn(config.RandomDelimiterCnt))
	delimiter := kvRepo.GetVal(ctx, delimiterKey)
	return luckName + delimiter + luckNum, nil
}
