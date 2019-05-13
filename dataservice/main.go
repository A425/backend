package main

import (
	"backend/dataservice/common"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"time"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(dalcommon.ServiceID),
		micro.Version(dalcommon.ServiceVersion),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*20),
	)

	// Initialise service
	service.Init()

	// Register Handler

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
