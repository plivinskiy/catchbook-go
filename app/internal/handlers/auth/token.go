package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) token(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := h.authUseCase.Authorize(username, password)
	if err != nil {
		message := fmt.Errorf("authorize error: %w", err)
		h.logger.Error(message.Error())
		c.JSON(http.StatusForbidden, gin.H{"message": message.Error()})
		c.Abort()
		return
	}
	tokens, err := h.authUseCase.Token(user)
	if err != nil {
		message := fmt.Errorf("authorize error: %w", err)
		h.logger.Error(message.Error())
		c.JSON(http.StatusForbidden, gin.H{"message": message.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, tokens)
}
