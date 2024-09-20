package project_routes

import (
	projectController "portfolioAPI/project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine){
  projectController := projectController.NewProjectController() 
  routerGroup := router.Group("/project", projectController.GetAllProjects)
  {
    routerGroup.GET("", )
  }
}
