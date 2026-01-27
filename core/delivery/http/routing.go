package http

import (
	"github.com/gin-gonic/gin"
)

type RoutingHandler interface {
	RegisterHandlers(routes *gin.RouterGroup)
	Handle() gin.HandlerFunc
}

func (h *routingHandler) RegisterHandlers(routes *gin.RouterGroup) {
	routes.POST("/auth/login", h.Handle())
	routes.POST("/auth/register", h.Handle())
}
