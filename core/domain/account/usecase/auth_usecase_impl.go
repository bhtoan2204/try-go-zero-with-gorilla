package usecase

import (
	"context"
	"errors"
	"fmt"
	appCtx "go-socket/core/context"
	"go-socket/core/domain/account/entity"
	accountrepos "go-socket/core/domain/account/repos"
	repos "go-socket/core/domain/account/repos"
	"go-socket/shared/infra/xpaseto"
	"go-socket/shared/pkg/hasher"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authUsecaseImpl struct {
	accountRepo accountrepos.AccountRepository
	hasher      hasher.Hasher
	paseto      xpaseto.PasetoService
}

func NewAuthUsecase(appCtx *appCtx.AppContext, repos repos.Repos) AuthUsecase {
	return &authUsecaseImpl{
		accountRepo: repos.AccountRepository(),
		hasher:      appCtx.GetHasher(),
		paseto:      appCtx.GetPaseto(),
	}
}

func (u *authUsecaseImpl) Login(ctx context.Context, email string, password string) (string, error) {
	account, err := u.accountRepo.GetAccountByEmail(ctx, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrAccountNotFound
	}
	valid, err := u.hasher.Verify(ctx, password, account.Password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", ErrInvalidCredentials
	}
	token, _, err := u.paseto.GenerateToken(ctx, account)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *authUsecaseImpl) Register(ctx context.Context, email string, password string) (string, error) {
	_, err := u.accountRepo.GetAccountByEmail(ctx, email)
	if err == nil {
		return "", ErrAccountExists
	}
	hashedPassword, err := u.hasher.Hash(ctx, password)
	if err != nil {
		return "", err
	}
	newAccountEntity := &entity.Account{
		ID:       uuid.New().String(),
		Email:    email,
		Password: hashedPassword,
	}
	if err := u.accountRepo.CreateAccount(ctx, newAccountEntity); err != nil {
		return "", fmt.Errorf("create account failed: %w", err)
	}
	token, _, err := u.paseto.GenerateToken(ctx, newAccountEntity)
	if err != nil {
		return "", fmt.Errorf("generate token failed: %w", err)
	}
	return token, nil
}

func (u *authUsecaseImpl) Logout(ctx context.Context, token string) error {
	return nil
}
