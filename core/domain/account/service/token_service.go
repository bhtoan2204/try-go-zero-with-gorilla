package service

import (
	"context"
	"time"

	"go-socket/core/domain/account/entity"
)

type TokenPayload struct {
	AccountID string
	Email     string
	ExpiresAt time.Time
}

type TokenService interface {
	GenerateToken(ctx context.Context, account *entity.Account) (string, time.Time, error)
	ParseToken(ctx context.Context, token string) (*TokenPayload, error)
}
