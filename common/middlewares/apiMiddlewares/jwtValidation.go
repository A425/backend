package apiMiddlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
	"net/http"
	"strings"
	"time"
)

const (
	authHeaderKey   = "Authorization"
	authTokenPrefix = "Bearer"
)

// JWTValidation check JWT token
func JWTValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(authHeaderKey)
		token = strings.Trim(token, authTokenPrefix)
		token = strings.TrimSpace(token)

		if token == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		t, err := jwt.Parse(token, keyFunc)
		if err != nil {
			log.Log("invalid token")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			userID := claims["userid"]
			if userIDStr, ok := userID.(string); !ok || userIDStr == "" {
				log.Logf("invalid userID: %s \n", userIDStr)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			expireAtRaw := claims["expireat"]

			expireAt, ok := expireAtRaw.(int64)
			if !ok {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			now := time.Now().Unix()
			if expireAt <= now {
				log.Log("expired jwt")
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			c.Next()
			return

		}

		log.Log("invalid claims")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte("hmacSampleSecret"), nil
}
