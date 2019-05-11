package apiscommon

import (
	"backend/common"
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-log/log"
)

// CreateCtxFromGinContext ...
func CreateCtxFromGinContext(c *gin.Context) (context.Context, error) {
	ctx := context.Background()

	t, err := common.GetJWTFromGinCtx(c)
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		userID := claims["userid"]
		userIDStr, ok := userID.(string)
		if !ok || userIDStr == "" {
			log.Logf("invalid userID: %s \n", userIDStr)
			return nil, errors.New("invalid userid")
		}

		ctx = context.WithValue(ctx, common.UserIDCtxKey, userIDStr)
	}

	return ctx, nil
}
