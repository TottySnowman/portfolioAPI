package tagRoutes

import (
	tagController "portfolioAPI/tag/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTagRoutes(router *gin.Engine){
  tagController := tagController.NewTagController()
  routerGroup := router.Group("/tag")
  {
    routerGroup.GET("/all", tagController.GetAllTags)
  }
}
