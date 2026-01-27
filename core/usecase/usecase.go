package usecase

import (
	accountusecase "go-socket/core/domain/account/usecase"
)

type AuthUsecase = accountusecase.AuthUsecase

type Usecase interface {
	AuthUsecase() accountusecase.AuthUsecase
}

type usecase struct {
	auth accountusecase.AuthUsecase
}

func NewUsecase(auth accountusecase.AuthUsecase) Usecase {
	return &usecase{
		auth: auth,
	}
}

func (u *usecase) AuthUsecase() accountusecase.AuthUsecase {
	return u.auth
}
