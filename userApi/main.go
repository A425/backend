package main

import (
	"net/http"
	"time"

	apimiddlewares "backend/common/middlewares/apiMiddlewares"
	"backend/userApi/client"
	userapicommon "backend/userApi/common"
	"backend/userApi/handler"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"golang.org/x/time/rate"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name(userapicommon.ServiceID),
		web.Version(userapicommon.ServiceVersion),
		web.RegisterInterval(time.Second*10),
		web.RegisterTTL(time.Second*20),
	)

	// Initialise service
	service.Init()

	client.InitAuthClient()

	// Register Handler
	router := gin.Default()
	h := new(handler.User)
	router.Use(gin.Recovery())

	apiGroupV1 := router.Group("/userapi/v1")

	{
		identifyGroup := apiGroupV1.Group("/identify")
		identifyRateLimiter := apimiddlewares.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP() // limit rate by client ip
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(time.Second), 10), time.Hour // limit 10 qps/clientIp, and the limiter liveness time duration is 1 hour
		}, func(c *gin.Context) {
			c.AbortWithStatus(http.StatusTooManyRequests) // handle exceed rate limit request
		})

		identifyGroup.Use(identifyRateLimiter)

		identifyGroup.POST("/wechat", h.LoginOrRegisterViaWechat)
		identifyGroup.POST("/refresh_token", h.RefreshAccessToken)
	}

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
