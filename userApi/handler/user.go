package handler

import (
	apicommon "backend/common/apis"
	userapicommon "backend/userApi/common"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"

	authCenterClient "backend/authCenter/proto/auth"
	"backend/userApi/client"

	"github.com/micro/go-micro/errors"
)

// User ...
type User struct{}

// WechatIdentifyReq 解析wechat请求
type WechatIdentifyReq struct {
	Code string `json:"code"`
}

// WechatIdentifyResp 返回wechat resp
type WechatIdentifyResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginOrRegisterViaWechat ...
func (e *User) LoginOrRegisterViaWechat(c *gin.Context) {
	fname := "LoginOrRegisterViaWechat"

	log.Logf("Received User.%s request", fname)

	reqBody := &WechatIdentifyReq{}
	err := c.BindJSON(reqBody)
	if err != nil {
		log.Logf(fname+": %+v ,err:%v \n", reqBody, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	cl := client.AuthClient()
	ctx := context.Background()
	// make request
	response, err := cl.VerifyWechatCode(ctx, &authCenterClient.VerifyWechatCodeRequest{
		Code: reqBody.Code,
	})
	if err != nil {
		apicommon.SendError(c, userapicommon.ServiceID+fname, err)
		return
	}

	fmt.Println(response.UnionID)

	// TODO Get or Create user in DB

	tokenResp, err := cl.GenerateTokens(ctx, &authCenterClient.GetTokensRequest{
		UserID: "user id from db",
	})

	if err != nil {
		apicommon.SendError(c, userapicommon.ServiceID+fname, err)
		return
	}

	rsp := WechatIdentifyResp{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
	}

	c.JSON(http.StatusOK, rsp)
}

// RefreshAccessTokenReq 刷新app的access token
type RefreshAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshAccessTokenResp 返回wechat resp
type RefreshAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshAccessToken ...
func (e *User) RefreshAccessToken(c *gin.Context) {
	fname := "RefreshAccessToken"

	log.Logf("Received User.%s request", fname)

	ctx, err := apicommon.CreateCtxFromGinContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest(userapicommon.ServiceID+fname, err.Error()))
		return
	}

	reqBody := &RefreshAccessTokenReq{}
	err = c.BindJSON(reqBody)
	if err != nil {
		log.Logf(fname+": %+v ,err:%v \n", reqBody, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	cl := client.AuthClient()

	tokenResp, err := cl.RefreshJWTToken(ctx, &authCenterClient.RefreshJWTRequest{
		RefreshToken: reqBody.RefreshToken,
	})

	if err != nil {
		apicommon.SendError(c, userapicommon.ServiceID+fname, err)
		return
	}

	rsp := WechatIdentifyResp{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
	}

	c.JSON(http.StatusOK, rsp)
}
