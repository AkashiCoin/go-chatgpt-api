package chatgpt

import (
	"fmt"
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api"
	"github.com/AkashiCoin/go-chatgpt-api/server/common"
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginInfo := &api.LoginInfo{}
	loginInfo.Username = c.PostForm("username")
	loginInfo.Password = c.PostForm("password")
	if loginInfo.Username == "" && loginInfo.Password == "" {
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": api.ParseUserInfoErrorMessage})
			return
		}
	}

	authenticator := api.NewAuthenticator(loginInfo.Username, loginInfo.Password, api.ProxyUrl)
	if err := authenticator.Begin(); err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"status": "error", "message": err.Details})
		return
	}
	c.Writer.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s; Secure; SameSite=Lax", common.SESSION_TOKEN_KEY, authenticator.Result.SessionToken))

	c.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"accessToken": authenticator.GetAccessToken(),
	})
}
