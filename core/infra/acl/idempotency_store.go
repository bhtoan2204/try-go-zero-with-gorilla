package acl

import (
	"context"
	"time"

	"go-socket/core/acl/idempotency"
	"go-socket/core/infra/cache"
)

const idempotencyPrefix = "idempotency:"

type RedisIdempotencyStore struct {
	cache cache.Cache
}

var _ idempotency.Store = (*RedisIdempotencyStore)(nil)

func NewRedisIdempotencyStore(cache cache.Cache) *RedisIdempotencyStore {
	return &RedisIdempotencyStore{cache: cache}
}

func (s *RedisIdempotencyStore) TryLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	if s == nil || s.cache == nil {
		return true, nil
	}
	seconds := int64(ttl.Seconds())
	if seconds <= 0 {
		seconds = 1
	}
	return s.cache.SetNX(ctx, idempotencyPrefix+key, seconds, "locked")
}

func (s *RedisIdempotencyStore) MarkDone(ctx context.Context, key string, ttl time.Duration) error {
	if s == nil || s.cache == nil {
		return nil
	}
	seconds := int64(ttl.Seconds())
	if seconds <= 0 {
		seconds = 1
	}
	return s.cache.SetValWithExp(ctx, idempotencyPrefix+key, "done", seconds)
}

func (s *RedisIdempotencyStore) Release(ctx context.Context, key string) error {
	if s == nil || s.cache == nil {
		return nil
	}
	return s.cache.Delete(ctx, idempotencyPrefix+key)
}
