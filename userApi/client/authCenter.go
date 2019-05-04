package client

import (
	authCenterClient "backend/authCenter/proto/auth"
	"github.com/micro/go-micro/client"
)

var (
	c authCenterClient.AuthCenterService
)

// AuthClient retrieves the client
func AuthClient() authCenterClient.AuthCenterService {
	return c
}

// InitAuthClient ...
func InitAuthClient() {
	c = authCenterClient.NewAuthCenterService("", client.DefaultClient)
}
