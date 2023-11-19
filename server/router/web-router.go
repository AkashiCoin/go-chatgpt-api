package router

import (
	"embed"
	"github.com/AkashiCoin/go-chatgpt-api/internal/utils"
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api"
	"github.com/AkashiCoin/go-chatgpt-api/server/common"
	"github.com/AkashiCoin/go-chatgpt-api/server/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetWebRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte) {
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.Cache())
	router.GET("/auth/login", func(c *gin.Context) {
		data, _ := buildFS.ReadFile("web/dist/auth/login/index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	router.Use(static.Serve("/", utils.EmbedFolder(buildFS, "web/dist")))
	router.NoRoute(func(c *gin.Context) {
		if _, err := c.Request.Cookie(common.SESSION_TOKEN_KEY); c.Request.Header.Get(api.AuthorizationHeader) != "" ||
			err == nil {
			api.Proxy(c)
			return
		}
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
