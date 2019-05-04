package client

import (
	"context"

	authCenterClient "backend/authCenter/proto/auth"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type authClient struct{}

// AuthClientFromContext retrieves the client from the Context
func AuthClientFromContext(ctx context.Context) (authCenterClient.AuthCenterService, bool) {
	c, ok := ctx.Value(authClient{}).(authCenterClient.AuthCenterService)
	return c, ok
}

// AuthClientWrapper returns a wrapper for the AutchClient
func AuthClientWrapper(service micro.Service) server.HandlerWrapper {
	ac := authCenterClient.NewAuthCenterService("", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, authClient{}, ac)
			return fn(ctx, req, rsp)
		}
	}
}
