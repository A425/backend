package handler

import (
	"context"
	"fmt"

	authcommon "backend/authCenter/common"
	auth "backend/authCenter/proto/auth"
	"backend/common"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
)

// WechatAuth ...
type AuthCenter struct{}

const (
	WechatValidationEndpoint = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type Code2SessionResp struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明 https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/union-id.html
	ErrCode    int64  `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}

// VerifyWechatCode 校验登录凭证
func (a *AuthCenter) VerifyWechatCode(ctx context.Context, req *auth.VerifyWechatCodeRequest, rsp *auth.VerifyWechatCodeResponse) error {
	log.Log("Received AuthCenter.VerifyWechatCode request:" + req.GetCode())
	url := fmt.Sprintf(WechatValidationEndpoint, "appid", "secret", req.Code)

	// using github.com/json-iterator/go to unmarshal
	st, r, err := fasthttp.Get(nil, url)
	if err != nil {
		return errors.InternalServerError(authcommon.ServiceID+"VerifyWechatCode", err.Error())
	}

	resp := &Code2SessionResp{}

	rsp.OpenID = resp.OpenID
	rsp.SessionKey = resp.SessionKey
	rsp.UnionID = resp.OpenID

	fmt.Println(rsp)

	/*
		VerifyWechatCode 里边做的事情: (最好改名)
			1.用openid和unionid检查当前用户是否存在,若存在返回userID
			2.if userID not empty then
				update session token of this userID
			3.if userID is empty then
				register 当前 openid 和 unionid 以及这个session key到数据库，并返回userID
			4.将 userid和一些时间拼成 jwt的入参并生成jwt
	*/

	log.Logf("st:%d, resp:%s", st, r)
	return nil
}

// GenerateTokens ...
func (a *AuthCenter) GenerateTokens(ctx context.Context, req *auth.GetTokensRequest, rsp *auth.GetTokensResponse) error {
	log.Log("Received AuthCenter.GetJWTToken request:" + req.GetUserID())

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := common.GenerateJWT(req.GetUserID())
	log.Log("tokenString:" + tokenString)
	if err != nil {
		return errors.InternalServerError(authcommon.ServiceID+"GenerateTokens", err.Error())
	}

	refreshToken := uuid.NewV4().String()

	// store refresh token to db

	rsp.AccessToken = tokenString
	rsp.RefreshToken = refreshToken

	return nil
}

// RefreshJWTToken ...
func (a *AuthCenter) RefreshJWTToken(ctx context.Context, req *auth.RefreshJWTRequest, rsp *auth.RefreshJWTResponse) error {
	log.Log("Received AuthCenter.RefreshJWTToken request:" + req.GetRefreshToken())

	userIDRaw := ctx.Value(common.UserIDCtxKey)
	userID, ok := userIDRaw.(string)
	if !ok {
		return errors.BadRequest(authcommon.ServiceID+"RefreshJWTToken", "invalid userid")
	}

	fmt.Println(userID)

	/*
		Get user from db
		Compare token between db data and input data

		if anything goes wrong then return with 403
	*/

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := common.GenerateJWT(userID)
	log.Log("tokenString:" + tokenString)
	if err != nil {
		return errors.InternalServerError(authcommon.ServiceID+"RefreshJWTToken", err.Error())
	}

	refreshToken := uuid.NewV4().String()

	// store refresh token to db

	rsp.AccessToken = tokenString
	rsp.RefreshToken = refreshToken

	return nil
}
