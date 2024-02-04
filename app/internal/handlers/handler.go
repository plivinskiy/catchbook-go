package handlers

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	Register(r *gin.Engine)
}
