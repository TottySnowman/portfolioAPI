package tagController

import (
	"net/http"
	tagModel "portfolioAPI/tag/models"
	tagService "portfolioAPI/tag/services"
	"strconv"

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

func (con *TagController) InsertTag(context *gin.Context) {
	var tag *tagModel.Tag

	if err := context.ShouldBindJSON(&tag); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	newTag, err := con.service.Insert(*tag)
  if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, newTag)
}

func (con *TagController) UpdateTag(context *gin.Context) {
	var tag *tagModel.Tag

	if err := context.ShouldBindJSON(&tag); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	if err := con.service.Update(*tag); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

func (con *TagController) DeleteTag(context *gin.Context) {
	tagID, err := strconv.Atoi(context.Param("ID"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := con.service.Delete(tagID); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

