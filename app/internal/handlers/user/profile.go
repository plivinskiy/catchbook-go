package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) profile(c *gin.Context) {
	id := c.Param("id")
	u, err := h.fetchUserUseCase.Fetch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}
