package token

import (
	"context"
	"fmt"
	"time"

	"go-socket/core/domain/account/entity"
	accountservice "go-socket/core/domain/account/service"

	"github.com/o1egl/paseto"
)

type PasetoService struct {
	paseto       *paseto.V2
	symmetricKey []byte
	issuer       string
	ttl          time.Duration
}

func NewPasetoService(symmetricKey string, issuer string, ttl time.Duration) (*PasetoService, error) {
	keyBytes := []byte(symmetricKey)
	if len(keyBytes) != 32 {
		return nil, fmt.Errorf("paseto key must be 32 bytes")
	}
	if ttl <= 0 {
		return nil, fmt.Errorf("token ttl must be positive")
	}
	return &PasetoService{
		paseto:       paseto.NewV2(),
		symmetricKey: keyBytes,
		issuer:       issuer,
		ttl:          ttl,
	}, nil
}

func (p *PasetoService) GenerateToken(ctx context.Context, account *entity.Account) (string, time.Time, error) {
	if account == nil {
		return "", time.Time{}, fmt.Errorf("account is nil")
	}
	now := time.Now().UTC()
	exp := now.Add(p.ttl)
	payload := paseto.JSONToken{
		Issuer:     p.issuer,
		Subject:    account.ID,
		IssuedAt:   now,
		Expiration: exp,
	}
	payload.Set("email", account.Email)

	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	if err != nil {
		return "", time.Time{}, err
	}
	return token, exp, nil
}

func (p *PasetoService) ParseToken(ctx context.Context, token string) (*accountservice.TokenPayload, error) {
	var jsonToken paseto.JSONToken
	var custom map[string]interface{}
	if err := p.paseto.Decrypt(token, p.symmetricKey, &jsonToken, &custom); err != nil {
		return nil, err
	}
	if !jsonToken.Expiration.IsZero() && time.Now().After(jsonToken.Expiration) {
		return nil, fmt.Errorf("token expired")
	}
	email, _ := custom["email"].(string)
	return &accountservice.TokenPayload{
		AccountID: jsonToken.Subject,
		Email:     email,
		ExpiresAt: jsonToken.Expiration,
	}, nil
}
