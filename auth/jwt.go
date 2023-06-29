package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

func Generate(userID string, roles []string, lifespan time.Duration, secret string) (string, error) {
	log.Debug().Str("user_id", userID).Msg("Generating token")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["roles"] = roles
	claims["exp"] = time.Now().Add(lifespan).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractID(c *gin.Context, secret string) (string, error) {
	log.Debug().Msg("Extracting user ID from token")

	token, err := extract(c, secret)
	if err != nil {
		return "0", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["user_id"].(string)
		if !ok {
			return "0", fmt.Errorf("invalid user_id: %v", claims["user_id"])
		}

		return uid, nil
	}

	return "0", nil
}

func ExtractRoles(c *gin.Context, secret string) (string, []string, error) {
	log.Debug().Msg("Extracting user ID and Roles from token")

	token, err := extract(c, secret)
	if err != nil {
		return "0", []string{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["user_id"].(string)
		if !ok {
			return "0", []string{}, fmt.Errorf("invalid user_id: %v", claims["user_id"])
		}

		roles, ok := claims["roles"].([]interface{})
		if !ok {
			return uid, []string{}, fmt.Errorf("invalid roles: %v", claims["roles"])
		}

		rolesStr := make([]string, len(roles))
		for i, v := range roles {
			rolesStr[i] = fmt.Sprint(v)
		}

		return uid, rolesStr, nil
	}

	return "0", []string{}, nil
}

func extract(c *gin.Context, secret string) (*jwt.Token, error) {

	tokenString := c.Query("token")
	if tokenString == "" {
		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenString = strings.Split(bearerToken, " ")[1]
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
