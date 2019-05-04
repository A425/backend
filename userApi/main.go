package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"

	"backend/userApi/client"
	"backend/userApi/handler"
	"github.com/micro/go-web"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name("go.micro.api.userApi"),
		web.Version("0.0.1"),
	)

	// Initialise service
	service.Init()

	client.InitAuthClient()

	// Register Handler
	router := gin.Default()
	h := new(handler.User)
	router.GET("/userApi/user/:code", h.VerifyWechatCode)

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
