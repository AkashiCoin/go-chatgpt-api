package router

import (
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api/chatgpt"
	"github.com/AkashiCoin/go-chatgpt-api/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetRelayRouter(router *gin.Engine) {
	router.POST("/api/auth/login", chatgpt.Login)
	router.POST("/backend-api/login", chatgpt.Login) // add support for other projects

	conversationGroup := router.Group("/backend-api/conversation", middleware.Authorization())
	{
		conversationGroup.POST("", chatgpt.CreateConversation)
	}
}
