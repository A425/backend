package apimiddlewares

import (
	"backend/common"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
)

// JWTValidation check JWT token
func JWTValidation() gin.HandlerFunc {
	return func(c *gin.Context) {

		t, err := common.GetJWTFromGinCtx(c)
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
