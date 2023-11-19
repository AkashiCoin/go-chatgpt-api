package router

import (
	"embed"
	"github.com/AkashiCoin/go-chatgpt-api/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte) {
	SetFreeCDNRouter(router)
	SetApiRouter(router)
	SetRelayRouter(router)
	router.Use(middleware.Session())
	SetWebRouter(router, buildFS, indexPage)
}
