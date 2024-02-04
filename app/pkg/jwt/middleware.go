package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c)
		if err != nil {
			unauthorised(c)
			return
		}
		fmt.Println("token:", string(token))
		c.Next()
	}
}

func unauthorised(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	c.Abort()
}

func extractToken(c *gin.Context) ([]byte, error) {
	authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		return nil, fmt.Errorf("token not found")
	}
	return []byte(authHeader[1]), nil
}
