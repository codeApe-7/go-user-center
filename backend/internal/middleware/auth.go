package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/codeApe-7/go-user-center/internal/utils"
)

type AuthConfig struct {
	JWTSecret string
}

func AuthRequired(cfg AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")
		claims, err := utils.ParseToken(cfg.JWTSecret, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("userUuid", claims.UserUUID)
		c.Set("email", claims.Email)
		c.Next()
	}
}
