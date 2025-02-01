package knowledgeRoutes

import (
	"github.com/gin-gonic/gin"
	knowledgeController "portfolioAPI/knowledge/controllers"
	middleware "portfolioAPI/router/authorization"
)

func RegisterKnowledgeRoutes(router *gin.Engine, knowledgeController *knowledgeController.KnowledgeController) {
	routerGroup := router.Group("/knowledge")
	{
		routerGroup.POST("", middleware.JWTMiddleware(), knowledgeController.UpsertText)
		routerGroup.PUT("", middleware.JWTMiddleware(), knowledgeController.UpsertText)
		routerGroup.DELETE("", middleware.JWTMiddleware(), knowledgeController.DeleteSinglePoint)
		routerGroup.GET("", middleware.JWTMiddleware(), knowledgeController.GetKnowledgeBase)
		routerGroup.POST("/fullSync", middleware.JWTMiddleware(), knowledgeController.FullSync)
		routerGroup.POST("/sync", middleware.JWTMiddleware(), knowledgeController.Sync)
	}
}
