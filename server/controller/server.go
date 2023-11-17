package controller

import (
	"github.com/AkashiCoin/gin-template/internal/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data": gin.H{
			"start_time": conf.StartTime,
			"app_name":   conf.AppName,
			"version":    conf.Version,
		},
	})
	return
}
