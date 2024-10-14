package project_routes

import (
	projectController "portfolioAPI/project/controllers"

	"github.com/gin-gonic/gin"
  middleware "portfolioAPI/router/authorization"
)

func RegisterProjectRoutes(router *gin.Engine){
  projectController := projectController.NewProjectController() 
  routerGroup := router.Group("/project")
  {
    routerGroup.GET("", projectController.GetAllProjects)
    routerGroup.POST("", middleware.JWTMiddleware(), projectController.InsertProject)
    routerGroup.PUT("", middleware.JWTMiddleware(), projectController.UpdateProject)
  }
}
