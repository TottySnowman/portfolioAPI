package chatRoutes

import (
	chatController "portfolioAPI/chat/controllers"
	middleware "portfolioAPI/router/authorization"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.Engine, chatController *chatController.ChatController) {
	routerGroup := router.Group("/chat")
	{
		routerGroup.POST("",middleware.JWTMiddleware(), chatController.Test)
	}
}

