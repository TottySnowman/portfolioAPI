package chatRoutes

import (
	"github.com/gin-gonic/gin"
	chatController "portfolioAPI/chat/controllers"
)

func RegisterChatRoutes(router *gin.Engine, chatController *chatController.ChatController) {
	routerGroup := router.Group("/chat")
	{
		routerGroup.POST("", chatController.Chat)
		routerGroup.GET("/ws", chatController.CreateWsConnection)
	}
}
