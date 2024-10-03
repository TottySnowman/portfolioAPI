package project_routes

import (
	projectController "portfolioAPI/project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine){
  projectController := projectController.NewProjectController() 
  routerGroup := router.Group("/project")
  {
    routerGroup.GET("", projectController.GetAllProjects)
    routerGroup.POST("", projectController.InsertProject)
  }
}
