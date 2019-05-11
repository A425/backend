package handler

import (
	"context"

	"github.com/micro/go-log"

	u "backend/dataservice/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (u *User) CreateUser(ctx context.Context, req *u.CreateUserReq, rsp *u.CreateUserResp) error {
	log.Log("Received User.CreateUser request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
