package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c)
		if err != nil {
			fmt.Println(token)
			c.Status(403)
			c.Abort()
			return
		}
		c.Next()
	}
}

func extractToken(c *gin.Context) ([]byte, error) {
	authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		return nil, fmt.Errorf("token not found")
	}
	return []byte(authHeader[1]), nil
}
