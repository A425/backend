package handler

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-log"

	auth "backend/authCenter/proto/auth"
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
		return err
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

	/*
		CheckSession 要做的: 当jwt过期或者验证失败时，
			1.
	*/

	log.Logf("st:%d, resp:%s", st, r)
	return nil
}

func (a *AuthCenter) GetJWTToken(ctx context.Context, req *auth.GetJWTTokenRequest, rsp *auth.GetJWTTokenResponse) error {
	log.Log("Received AuthCenter.GetJWTToken request:" + req.GetUserID())

	// TODO Check userID

	now := time.Now()
	expireAt := now.Add(time.Hour * 2)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   req.GetUserID(),
		"issueat":  now.Unix(),      // 签发时间
		"expireat": expireAt.Unix(), // 过期时间
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("hmacSampleSecret"))

	fmt.Println(tokenString, err)

	return err
}
