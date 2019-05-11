package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey   = "Authorization"
	authTokenPrefix = "Bearer"
)

// GetJWTFromGinCtx ...
func GetJWTFromGinCtx(c *gin.Context) (*jwt.Token, error) {
	token := c.GetHeader(authHeaderKey)
	token = strings.Trim(token, authTokenPrefix)
	token = strings.TrimSpace(token)

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("hmacSampleSecret"), nil
	})
}

// GenerateJWT 生成 jwt
func GenerateJWT(userID string) (string, error) {
	now := time.Now()
	expireAt := now.Add(time.Minute * 10)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userID,
		"issueat":  now.Unix(),      // 签发时间
		"expireat": expireAt.Unix(), // 过期时间
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("hmacSampleSecret"))

	return tokenString, err
}
