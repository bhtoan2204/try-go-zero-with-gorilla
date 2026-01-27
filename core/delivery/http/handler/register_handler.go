package handler

import (
	"errors"
	"go-socket/core/delivery/http/data/in"
	"go-socket/core/delivery/http/data/out"
	"go-socket/core/usecase"
	"go-socket/shared/pkg/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type registerHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewRegisterHandler(usecase usecase.Usecase) RequestHandler {
	return &registerHandler{
		authUsecase: usecase.AuthUsecase(),
	}
}

func (h *registerHandler) Handle(c *gin.Context) (interface{}, error) {
	ctx := c.Request.Context()
	logger := logging.FromContext(ctx)
	var request in.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Errorw("Unmarshal request failed", zap.Error(err))
		return nil, err
	}
	if err := request.Validate(); err != nil {
		logger.Errorw("Validate request failed", zap.Error(err))
		return nil, errors.New("validate request failed")
	}
	token, err := h.authUsecase.Register(ctx, request.Email, request.Password)
	if err != nil {
		logger.Errorw("Register failed", zap.Error(err))
		return nil, errors.New("register failed")
	}
	return out.RegisterResponse{
		Token: token,
	}, nil
}
