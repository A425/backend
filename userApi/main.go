package main

import (
	"github.com/micro/go-log"

	"backend/userApi/client"
	"backend/userApi/handler"
	"github.com/micro/go-micro"

	user "backend/userApi/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.userApi"),
		micro.Version("0.0.1"),
	)

	// Initialise service
	service.Init(
		// create wrap for the AuthClient srv client
		micro.WrapHandler(client.AuthClientWrapper(service)),
	)

	// Register Handler
	svr := service.Server()
	user.RegisterUserHandler(svr, new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
