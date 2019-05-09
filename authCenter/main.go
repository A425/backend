package main

import (
	"backend/authCenter/handler"
	authProto "backend/authCenter/proto/auth"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.authCenter"),
		micro.Version("0.0.1"),
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
