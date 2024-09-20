package projectController

import (
	"net/http"
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

