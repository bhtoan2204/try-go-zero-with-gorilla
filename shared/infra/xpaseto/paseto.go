package xpaseto

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"go-socket/core/domain/account/entity"

	"github.com/o1egl/paseto"
)

type PasetoPayload struct {
	AccountID string
	Email     string
	ExpiresAt time.Time
}

type PasetoService interface {
	GenerateToken(ctx context.Context, account *entity.Account) (string, time.Time, error)
	ParseToken(ctx context.Context, token string) (*PasetoPayload, error)
}

type pasetoService struct {
	paseto       *paseto.V2
	symmetricKey []byte
	issuer       string
	ttl          time.Duration
}

func NewPaseto(symmetricKey string, issuer string, ttl time.Duration) (PasetoService, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(symmetricKey)
	if err != nil {
		return nil, err
	}
	if len(keyBytes) != 32 {
		return nil, fmt.Errorf("paseto key must be 32 bytes")
	}
	if ttl <= 0 {
		return nil, fmt.Errorf("token ttl must be positive")
	}
	return &pasetoService{
		paseto:       paseto.NewV2(),
		symmetricKey: keyBytes,
		issuer:       issuer,
		ttl:          ttl,
	}, nil
}

func (p *pasetoService) GenerateToken(ctx context.Context, account *entity.Account) (string, time.Time, error) {
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

func (p *pasetoService) ParseToken(ctx context.Context, token string) (*PasetoPayload, error) {
	var jsonToken paseto.JSONToken
	var custom map[string]interface{}
	if err := p.paseto.Decrypt(token, p.symmetricKey, &jsonToken, &custom); err != nil {
		return nil, err
	}
	if !jsonToken.Expiration.IsZero() && time.Now().After(jsonToken.Expiration) {
		return nil, fmt.Errorf("token expired")
	}
	email, _ := custom["email"].(string)
	return &PasetoPayload{
		AccountID: jsonToken.Subject,
		Email:     email,
		ExpiresAt: jsonToken.Expiration,
	}, nil
}
