package tagRoutes

import (
	middleware "portfolioAPI/router/authorization"
	tagController "portfolioAPI/tag/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTagRoutes(router *gin.Engine, tagController *tagController.TagController) {
	routerGroup := router.Group("/tag")
	{
		routerGroup.GET("",middleware.JWTMiddleware(), tagController.GetAllTags)
		routerGroup.POST("", middleware.JWTMiddleware(), tagController.InsertTag)
		routerGroup.PUT("", middleware.JWTMiddleware(), tagController.UpdateTag)
		routerGroup.DELETE("/:ID", middleware.JWTMiddleware(), tagController.DeleteTag)
	}
}
