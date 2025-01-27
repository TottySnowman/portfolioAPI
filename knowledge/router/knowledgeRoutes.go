package knowledgeRoutes

import (
	knowledgeController "portfolioAPI/knowledge/controllers"
	middleware "portfolioAPI/router/authorization"
	"github.com/gin-gonic/gin"
)

func RegisterKnowledgeRoutes(router *gin.Engine, knowledgeController *knowledgeController.KnowledgeController) {
	routerGroup := router.Group("/knowledge")
	{
		routerGroup.POST("",middleware.JWTMiddleware(), knowledgeController.UpsertText)
		routerGroup.PUT("",middleware.JWTMiddleware(), knowledgeController.UpsertText)
		routerGroup.DELETE("",middleware.JWTMiddleware(), knowledgeController.UpsertText)
		routerGroup.GET("",middleware.JWTMiddleware(), knowledgeController.UpsertText)
	}
}
