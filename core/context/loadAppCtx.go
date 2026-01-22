package appCtx

import (
	"context"
	"go-socket/config"
	"go-socket/constant"
	"go-socket/core/infra/cache"
	"go-socket/core/infra/persistent"
	"go-socket/core/infra/redis"
)

func LoadAppCtx(ctx context.Context, cfg *config.Config) (*AppContext, error) {
	var opts []Option

	db, err := persistent.NewConnection(ctx, cfg)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithDB(db))

	redisClient, err := redis.NewStandaloneRedisClient(cfg)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithRedisClient(redisClient))

	cache := cache.New(redisClient, constant.DEFAULT_CACHE_EXPIRATION_TIME)
	opts = append(opts, WithCache(cache))

	return NewAppContext(ctx, opts...)
}
