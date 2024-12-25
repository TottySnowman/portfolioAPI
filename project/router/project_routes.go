package project_routes

import (
	projectController "portfolioAPI/project/controllers"

	"github.com/gin-gonic/gin"
	middleware "portfolioAPI/router/authorization"
)

func RegisterProjectRoutes(router *gin.Engine, projectController *projectController.ProjectController) {
	routerGroup := router.Group("/project")
	{
		routerGroup.GET("", projectController.GetAllProjects)
		routerGroup.GET("/all",middleware.JWTMiddleware(), projectController.GetAllProjectsIncludeHidden)
		routerGroup.POST("", middleware.JWTMiddleware(), projectController.InsertProject)
		routerGroup.PUT("", middleware.JWTMiddleware(), projectController.UpdateProject)
		routerGroup.DELETE("/:ID", middleware.JWTMiddleware(), projectController.DeleteProject)
	}
}
