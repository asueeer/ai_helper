package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	RedisHost       = "47.104.186.111:6379"
	Pwd             = "nearby123"
	DefaultDatabase = 0
	DatabaseForWx   = 1
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     RedisHost,
		Password: Pwd,
		DB:       DefaultDatabase, // use default DB
	})
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	return redisClient.Get(ctx, key)
}

func Set(ctx context.Context, key string, val string, expireAt time.Duration) error {
	statusCmd := redisClient.Set(ctx, key, val, expireAt)
	return statusCmd.Err()
}

func Push(ctx context.Context, key string, val string) error {
	return redisClient.LPush(ctx, key, val).Err()
}

func Pop(ctx context.Context, key string) *redis.StringCmd {
	return redisClient.LPop(ctx, key)
}

func PopCnt(ctx context.Context, key string, cnt int) *redis.StringSliceCmd {
	return redisClient.LPopCount(ctx, key, cnt)
}

func HIncrBy(ctx context.Context, key string, field string, incr int64) *redis.IntCmd {
	return redisClient.HIncrBy(ctx, key, field, incr)
}

func HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return redisClient.HScan(ctx, key, cursor, match, count)
}

func HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return redisClient.HDel(ctx, key, fields...)
}

func HSet(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return redisClient.HSet(ctx, key, fields)
}

func HMSet(ctx context.Context, key string, values ...string) *redis.BoolCmd {
	return redisClient.HMSet(ctx, key, values)
}

func HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	return redisClient.HGetAll(ctx, key)
}

func ExpireAt(ctx context.Context, key string, tm time.Time) {
	redisClient.ExpireAt(ctx, key, tm)
}

func SAdd(ctx context.Context, key string, members ...string) *redis.IntCmd {
	return redisClient.SAdd(ctx, key, members)
}

func SRemove(ctx context.Context, key string, members ...string) *redis.IntCmd {
	return redisClient.SRem(ctx, key, members)
}
