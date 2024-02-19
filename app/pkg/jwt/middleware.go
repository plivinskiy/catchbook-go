package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Middleware https://github.com/theartofdevel/notes_system/blob/main/api_service/app/pkg/jwt/middleware.go
func Middleware(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c)
		if err != nil {
			unauthorised(c)
			return
		}
		verifier, err := jwt.NewVerifierHS(jwt.HS256, secret)
		if err != nil {
			unauthorised(c)
			return
		}
		newToken, err := jwt.ParseAndVerify(token, verifier)
		if err != nil {
			unauthorised(c)
			return
		}
		var newClaims UserClaims
		errClaims := json.Unmarshal(newToken.RawClaims(), &newClaims)
		if errClaims != nil {
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
