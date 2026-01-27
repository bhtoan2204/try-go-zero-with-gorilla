package handler

import "github.com/gin-gonic/gin"

type RequestHandler interface {
	Handle(c *gin.Context) (interface{}, error)
}
