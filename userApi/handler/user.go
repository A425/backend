package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-log"

	authCenterClient "backend/authCenter/proto/auth"
	"backend/userApi/client"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

func (e *User) VerifyWechatCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.VerifyWechatCode request, path:" + req.GetPath())

	// extract the client from the context
	c, ok := client.AuthClientFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.userApi.user.VerifyWechatCode", "auth client not found")
	}

	// make request
	response, err := c.VerifyWechatCode(ctx, &authCenterClient.VerifyWechatCodeRequest{
		Code: extractValue(req.Post["code"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.userApi.user.VerifyWechatCode", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
