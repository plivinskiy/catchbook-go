package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Params struct {
	ID uint `uri:"id"`
}

func (h *Handler) profile(c *gin.Context) {
	var p Params
	err := c.ShouldBindUri(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong parameters"})
		h.logger.Error(err.Error())
		return
	}
	u, err := h.fetchUserUseCase.FetchUser(p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not found"})
		h.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, u)
}
