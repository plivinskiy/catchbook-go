package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	list, err := h.listUserUseCase.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}
