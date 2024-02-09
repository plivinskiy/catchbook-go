package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) token(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := h.authUseCase.Authorize(username, password)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"message": "login or password wrong"})
		c.Abort()
		return
	}
	tokens, err := h.authUseCase.Token(user)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"message": "cannot generate token"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, tokens)
}
