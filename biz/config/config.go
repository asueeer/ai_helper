package config

import (
	"ai_helper/biz/dal/db/repo"
	"context"
	"log"
)

var (
	WorkID = int64(1)

	// DefaultHeadURLCnt 库里存了一些随机头像、随机昵称、随机数字、随机分隔符
	DefaultHeadURLCnt  = -1
	LuckyNameCnt       = -1
	LuckyNumberCnt     = -1
	RandomDelimiterCnt = -1
)

// InitKVConfig
// 依赖数据库
// 需要数据库初始化成功之后再进行调用
func InitKVConfig() {
	var err error
	ctx := context.Background()
	bizKVRepo := repo.NewBizKVRepo(ctx)

	/*
		对统计值进行初始化
		方便生成随机值，
	*/
	{
		DefaultHeadURLCnt, err = bizKVRepo.GetDefaultHeadURLCnt(ctx)
		if err != nil {
			panic(err)
		}
		LuckyNameCnt, err = bizKVRepo.GetLuckyNameCnt(ctx)
		if err != nil {
			panic(err)
		}
		LuckyNumberCnt, err = bizKVRepo.GetLuckyNumberCnt(ctx)
		if err != nil {
			panic(err)
		}
		RandomDelimiterCnt, err = bizKVRepo.GetRandomDelimiterCnt(ctx)
		if err != nil {
			panic(err)
		}

		log.Printf("DefaultHeadURLCnt: %+v", DefaultHeadURLCnt)
		log.Printf("LuckyNameCnt: %+v", LuckyNameCnt)
		log.Printf("LuckyNumberCnt: %+v", LuckyNumberCnt)
		log.Printf("RandomDelimiterCnt: %+v", RandomDelimiterCnt)
	}
}
