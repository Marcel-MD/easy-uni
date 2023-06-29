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

func JwtAuthRoles(secret string, requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, roles, err := auth.ExtractRoles(c, secret)
		if err != nil {
			log.Err(err).Msg("Invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if !contains(roles, requiredRoles) {
			log.Err(err).Msg("Invalid roles")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", id)
		c.Set("roles", roles)
		c.Next()
	}
}

func contains(s []string, e []string) bool {
	for _, a := range e {
		for _, b := range s {
			if b == a {
				return true
			}
		}
	}
	return false
}
