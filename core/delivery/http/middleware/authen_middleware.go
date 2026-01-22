package middleware

import "github.com/gin-gonic/gin"

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
