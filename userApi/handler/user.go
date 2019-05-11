package handler

import (
	apicommon "backend/common/apis"
	userapicommon "backend/userApi/common"
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
	log.Logf("Received User.LoginOrRegisterViaWechat request, d")

	cl := client.AuthClient()

	ctx, err := apicommon.CreateCtxFromGinContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest(userapicommon.ServiceID+"LoginOrRegisterViaWechat", err.Error()))
		return
	}

	reqBody := &WechatIdentifyReq{}
	err = c.BindJSON(reqBody)
	if err != nil {
		log.Logf("LoginOrRegisterViaWechat: %+v ,err:%v \n", reqBody, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// make request
	response, err := cl.VerifyWechatCode(ctx, &authCenterClient.VerifyWechatCodeRequest{
		Code: reqBody.Code,
	})
	if err != nil {
		c.JSON(500, errors.InternalServerError(userapicommon.ServiceID+"LoginOrRegisterViaWechat", err.Error()))
		return
	}

	// TODO Get or Create user in DB

	cl.GenerateTokens(ctx, &authCenterClient.GetTokensRequest{
		UserID: "user id from db",
	})

	c.JSON(200, response)
}
