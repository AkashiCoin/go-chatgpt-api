package router

import (
	"github.com/AkashiCoin/go-chatgpt-api/server/controller"
	"github.com/AkashiCoin/go-chatgpt-api/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middleware.CORS())
	{
		api.GET("/status", controller.GetStatus)
	}
}
