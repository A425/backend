package client

import (
	"context"

	authCenterClient "backend/authCenter/proto/example"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type exampleKey struct{}

// FromContext retrieves the client from the Context
func ExampleFromContext(ctx context.Context) (authCenterClient.ExampleService, bool) {
	c, ok := ctx.Value(exampleKey{}).(authCenterClient.ExampleService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func ExampleWrapper(service micro.Service) server.HandlerWrapper {
	client := authCenterClient.NewExampleService("", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
