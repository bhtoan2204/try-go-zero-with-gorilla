package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go-socket/core/domain/account/entity"
	accountrepos "go-socket/core/domain/account/repos"
	accountservice "go-socket/core/domain/account/service"
	"go-socket/core/infra/hasher"

	"github.com/google/uuid"
)

type authUsecase struct {
	accountRepo  accountrepos.AccountRepository
	tokenRepo    accountrepos.TokenRepository
	hasher       hasher.Hasher
	tokenService accountservice.TokenService
}

func NewAuthUsecase(
	accountRepo accountrepos.AccountRepository,
	tokenRepo accountrepos.TokenRepository,
	hasher hasher.Hasher,
	tokenService accountservice.TokenService,
) AuthUsecase {
	return &authUsecase{
		accountRepo:  accountRepo,
		tokenRepo:    tokenRepo,
		hasher:       hasher,
		tokenService: tokenService,
	}
}

func (u *authUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return "", ErrInvalidCredentials
	}
	account, err := u.accountRepo.GetAccountByEmail(ctx, email)
	if err != nil {
		if err == accountrepos.ErrAccountNotFound {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	ok, err := u.hasher.Verify(ctx, password, account.Password)
	if err != nil || !ok {
		return "", ErrInvalidCredentials
	}
	token, _, err := u.tokenService.GenerateToken(ctx, account)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUsecase) Register(ctx context.Context, email string, password string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return "", fmt.Errorf("email and password are required")
	}
	if _, err := u.accountRepo.GetAccountByEmail(ctx, email); err == nil {
		return "", ErrAccountExists
	} else if err != accountrepos.ErrAccountNotFound {
		return "", err
	}

	hashed, err := u.hasher.Hash(ctx, password)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	account := &entity.Account{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  hashed,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := u.accountRepo.CreateAccount(ctx, account); err != nil {
		return "", err
	}

	token, _, err := u.tokenService.GenerateToken(ctx, account)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUsecase) Logout(ctx context.Context, token string) error {
	if strings.TrimSpace(token) == "" {
		return nil
	}
	payload, err := u.tokenService.ParseToken(ctx, token)
	if err != nil {
		return err
	}
	return u.tokenRepo.RevokeToken(ctx, token, payload.ExpiresAt)
}
