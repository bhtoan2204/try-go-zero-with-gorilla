package handler

import (
	"context"
	"go-socket/core/pkg/logging"
	"go-socket/core/usecase"

	"github.com/gin-gonic/gin"
)

type loginHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewLoginHandler() RequestHandler {
	return &loginHandler{}
}

func (h *loginHandler) Handle(ctx context.Context) (interface{}, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("login handler")
	return gin.H{
		"message": "login handler",
	}, nil
}
