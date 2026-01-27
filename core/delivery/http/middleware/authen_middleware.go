package middleware

import (
	appCtx "go-socket/core/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware(appCtx *appCtx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
