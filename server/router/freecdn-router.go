package router

import (
	"github.com/AkashiCoin/go-chatgpt-api/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetFreeCDNRouter(router *gin.Engine) {
	router.Use(middleware.FreeCDNMiddleware())

	router.GET("/freecdn-boot", middleware.FreeCDNBootHandler)
	router.GET("/freecdn-nojs", middleware.FreeCDNNoJSHandler)
	router.GET("/freecdn-goto", middleware.FreeCDNGotoHandler)
}
