package cache

import (
	"context"
	"time"

	"go-socket/core/domain/account/repos"
)

const tokenBlacklistPrefix = "token:blacklist:"

type tokenRepo struct {
	cache Cache
}

var _ repos.TokenRepository = (*tokenRepo)(nil)

func NewTokenRepo(cache Cache) repos.TokenRepository {
	return &tokenRepo{cache: cache}
}

func (r *tokenRepo) RevokeToken(ctx context.Context, token string, expiresAt time.Time) error {
	if r == nil || r.cache == nil {
		return nil
	}
	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		return nil
	}
	seconds := int64(ttl.Seconds())
	if seconds <= 0 {
		return nil
	}
	return r.cache.SetValWithExp(ctx, tokenBlacklistPrefix+token, "revoked", seconds)
}

func (r *tokenRepo) IsTokenRevoked(ctx context.Context, token string) (bool, error) {
	if r == nil || r.cache == nil {
		return false, nil
	}
	if r.cache.Exists(ctx, tokenBlacklistPrefix+token) > 0 {
		return true, nil
	}
	return false, nil
}
