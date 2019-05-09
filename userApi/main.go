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
		web.Name("go.micro.api.userapi"),
		web.Version("0.0.1"),
	)

	// Initialise service
	service.Init()

	client.InitAuthClient()

	// Register Handler
	router := gin.Default()
	h := new(handler.User)
	router.Use(gin.Recovery())

	apiGroup := router.Group("/userapi")
	apiGroup.GET("/v1/login/wechat/:code", h.LoginOrRegisterViaWechat)

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
