package projectController

import (
	"net/http"
	projectModel "portfolioAPI/project/models"
	projectService "portfolioAPI/project/services"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
  service *projectService.ProjectService
}

func NewProjectController() *ProjectController{
  return &ProjectController{
    service: projectService.NewProjectService(),
  }
}

func (con *ProjectController) GetAllProjects(context *gin.Context){
 projects := con.service.GetAllProjects() 
  if projects == nil{
    context.JSON(http.StatusNotFound, gin.H{"error": "No projects found"})
    return
  }

  context.IndentedJSON(http.StatusOK, projects)
}

func (con *ProjectController) InsertProject(context *gin.Context){
  var project *projectModel.Project

  if err := context.ShouldBindJSON(&project); err != nil{
    context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
    return
  }

  if err := con.service.Insert(*project); err != nil{
    context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
    return
  }
  
  context.Status(http.StatusCreated)
}

func (con *ProjectController) UpdateProject(context *gin.Context){
  var project *projectModel.Project

  if err := context.ShouldBindJSON(&project); err != nil{
    context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
    return
  }

  if err := con.service.Update(*project); err != nil{
    context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
    return
  }
  
  context.Status(http.StatusOK)
}
