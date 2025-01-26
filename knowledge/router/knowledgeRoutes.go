package knowledgeRoutes

import (
	knowledgeController "portfolioAPI/knowledge/controllers"
	middleware "portfolioAPI/router/authorization"
	"github.com/gin-gonic/gin"
)

func RegisterKnowledgeRoutes(router *gin.Engine, knowledgeController *knowledgeController.KnowledgeController) {
	routerGroup := router.Group("/knowledge")
	{
		routerGroup.POST("/",middleware.JWTMiddleware(), knowledgeController.InsertText)
	}
}
