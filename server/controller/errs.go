package controller

import (
	"github.com/AkashiCoin/go-chatgpt-api/server/common"
	"github.com/gin-gonic/gin"
)

func ApiNotFound(c *gin.Context) {
	common.ErrorStrResp(c, "api not found", 404)
}
