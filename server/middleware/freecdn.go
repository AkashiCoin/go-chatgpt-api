package middleware

import (
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func FreeCDNMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("freecdn"); err == nil && cookie == "0" {
			c.Next()
		} else if c.GetHeader("Sec-Fetch-Dest") != "document" ||
			c.GetHeader("Pragma") == "no-cache" ||
			c.Request.Method != "GET" ||
			strings.HasPrefix(c.Request.URL.Path, "/freecdn-") ||
			strings.Contains(c.Request.URL.RawQuery, "freecdn__=0") ||
			strings.Contains(c.GetHeader("Referer"), "freecdn__=0") {
			c.Next()
		} else {
			FreeCDNBootHandler(c)
			c.Abort()
		}
	}
}

func FreeCDNBootHandler(c *gin.Context) {
	c.Header("Expires", "-1")
	c.Data(http.StatusOK, "text/html", []byte("<script src=/freecdn-loader.min.js></script><noscript><meta http-equiv=Refresh content=0;url=/freecdn-nojs></noscript>"))
}

func FreeCDNNoJSHandler(c *gin.Context) {
	c.SetCookie("freecdn", "0", 0, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/freecdn-goto?"+c.GetHeader("Referer"))
}

func FreeCDNGotoHandler(c *gin.Context) {
	c.Header("Expires", "-1")
	returnURL, _ := url.Parse(c.Query("return_url"))
	if returnURL == nil || returnURL.Host != c.Request.Host {
		c.String(http.StatusBadRequest, "invalid url or host")
		return
	}
	if cookie, err := c.Cookie("freecdn"); err == nil && cookie != "0" {
		query := returnURL.Query()
		query.Add("freecdn__", "0")
		returnURL.RawQuery = query.Encode()
	}
	c.Redirect(http.StatusTemporaryRedirect, returnURL.String())
}
