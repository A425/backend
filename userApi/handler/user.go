package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"

	authCenterClient "backend/authCenter/proto/auth"
	"backend/userApi/client"
	"github.com/micro/go-micro/errors"
)

type User struct{}

func ctxFromGinContext(c *gin.Context) context.Context {
	return context.Background()
}

func (e *User) LoginOrRegisterViaWechat(c *gin.Context) {
	log.Logf("Received User.LoginOrRegisterViaWechat request, d")

	cl := client.AuthClient()

	ctx := ctxFromGinContext(c)
	// make request
	response, err := cl.VerifyWechatCode(ctx, &authCenterClient.VerifyWechatCodeRequest{
		Code: c.Param("code"),
	})
	if err != nil {
		c.JSON(500, errors.InternalServerError("go.micro.api.userApi.user.VerifyWechatCode", err.Error()))
		return
	}

	// TODO Get or Create user in DB

	cl.GetJWTToken(ctx, &authCenterClient.GetJWTTokenRequest{
		// UserID: respon
	})

	c.JSON(200, response)
}
