package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-log"
	"github.com/valyala/fasthttp"

	wechatAuth "backend/authCenter/proto/auth"
)

// WechatAuth ...
type WechatAuth struct{}

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
func (a *WechatAuth) VerifyWechatCode(ctx context.Context, req *wechatAuth.VerifyWechatCodeRequest, rsp *wechatAuth.VerifyWechatCodeResponse) error {
	log.Log("Received WechatAuth.VerifyWechatCode request:" + req.GetCode())
	rsp.Msg = "Hello " + req.Code
	url := fmt.Sprintf(WechatValidationEndpoint, "appid", "secret", req.Code)

	// using github.com/json-iterator/go to unmarshal
	st, r, err := fasthttp.Get(nil, url)
	if err != nil {
		return err
	}

	resp := &Code2SessionResp{}

	fmt.Println(resp)

	/*
		VerifyWechatCode 里边做的事情: (最好改名)
			1.用openid和unionid检查当前用户是否存在,若存在返回userID
			2.if userID not empty then
				update session token of this userID
			3.if userID is empty then
				register 当前 openid 和 unionid 以及这个session key到数据库，并返回userID
			4.为当前用户生成一个uuid随机数存到dal中
			5.将 userid, 随机数和一些时间拼成 jwt的入参并生成jwt
	*/

	/*
		CheckSession 要做的: 当jwt过期或者验证失败时，
			1.
	*/

	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":      "userid",
		"uid":         "uuid",
		"issueat":     now,        // 签发时间
		"deprecateat": now + 5*60, // 过期时间
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString("hmacSampleSecret")

	fmt.Println(tokenString, err)

	log.Logf("st:%d, resp:%s", st, r)
	return nil
}
