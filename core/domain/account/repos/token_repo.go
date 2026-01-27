package repos

import (
	"context"
	"time"
)

type TokenRepository interface {
	RevokeToken(ctx context.Context, token string, expiresAt time.Time) error
	IsTokenRevoked(ctx context.Context, token string) (bool, error)
}
