package router

import (
	"github.com/AkashiCoin/gin-template/server/controller"
	"github.com/AkashiCoin/gin-template/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middleware.CORS())
	{
		api.GET("/status", controller.GetStatus)
	}
}
