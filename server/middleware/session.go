package middleware

import (
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api"
	"github.com/AkashiCoin/go-chatgpt-api/server/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Request.Cookie(common.SESSION_TOKEN_KEY); err == nil &&
			!strings.Contains(c.Request.URL.Path, "freecdn") &&
			strings.Contains(c.Request.RequestURI, "api") ||
			c.Request.Header.Get(api.AuthorizationHeader) != "" {
			api.Proxy(c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
