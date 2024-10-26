package tagController

import (
	"net/http"
	tagService "portfolioAPI/tag/services"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	service *tagService.TagService
}

func NewTagController() *TagController {
	return &TagController{
		service: tagService.NewTagService(),
	}
}

func(controller *TagController) GetAllTags(context *gin.Context){
  existingTags := controller.service.GetAllTags()

  if existingTags == nil{
    		context.JSON(http.StatusNotFound, gin.H{"error": "No tags found"})
		return
  }

  context.IndentedJSON(http.StatusOK, existingTags)
}
