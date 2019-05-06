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

// VerifyWechatCode 校验登录凭证
func (a *WechatAuth) VerifyWechatCode(ctx context.Context, req *wechatAuth.VerifyWechatCodeRequest, rsp *wechatAuth.VerifyWechatCodeResponse) error {
	log.Log("Received WechatAuth.VerifyWechatCode request:" + req.GetCode())
	rsp.Msg = "Hello " + req.Code
	url := fmt.Sprintf(WechatValidationEndpoint, "appid", "secret", req.Code)

	// using github.com/json-iterator/go to unmarshal
	st, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString("hmacSampleSecret")

	fmt.Println(tokenString, err)

	log.Logf("st:%d, resp:%s", st, resp)
	return nil
}
