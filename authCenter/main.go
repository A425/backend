package main

import (
	authcommon "backend/authCenter/common"
	"backend/authCenter/handler"
	authProto "backend/authCenter/proto/auth"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(authcommon.ServiceID),
		micro.Version(authcommon.ServiceVersion),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*20),
	)

	// Initialise service
	service.Init()

	// Register Handler
	svr := service.Server()
	authProto.RegisterAuthCenterHandler(svr, new(handler.AuthCenter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
