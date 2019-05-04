package handler

import (
	"context"

	"github.com/micro/go-log"

	wechatAuth "backend/authCenter/proto/auth"
)

// WechatAuth ...
type WechatAuth struct{}

// VerifyWechatCode 校验登录凭证
func (a *WechatAuth) VerifyWechatCode(ctx context.Context, req *wechatAuth.VerifyWechatCodeRequest, rsp *wechatAuth.VerifyWechatCodeResponse) error {
	log.Log("Received WechatAuth.VerifyWechatCode request:" + req.GetCode())
	rsp.Msg = "Hello " + req.Code
	return nil
}
