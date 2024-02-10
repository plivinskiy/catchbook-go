package user

import (
	"catchbook/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	var dto model.UserDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		h.logger.Error("wrong parameters: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong parameters"})
		return
	}
	u, err := h.createUserUseCase.CreateUser(dto)
	if err != nil {
		h.logger.Error("cannot save user: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot save user"})
		return
	}
	c.JSON(http.StatusOK, u)
}
