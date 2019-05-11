package main

import (
	"backend/dataservice/handler"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	example "backend/dataservice/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(dalcommon.ServiceID),
		micro.Version(dalcommon.ServiceVersion),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
