package middleware

import (
	"net/http"

	"github.com/Marcel-MD/easy-uni/auth"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := auth.ExtractID(c, secret)
		if err != nil {
			log.Err(err).Msg("Invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}
