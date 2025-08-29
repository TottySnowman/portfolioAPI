package projectController

import (
	"net/http"
	projectModel "portfolioAPI/project/models"
	projectService "portfolioAPI/project/services"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectService *projectService.ProjectService
}

func NewProjectController(projectService *projectService.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (con *ProjectController) GetAllProjects(context *gin.Context) {
	projects := con.projectService.GetAllProjects(false)
	if projects == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No projects found"})
		return
	}

	context.IndentedJSON(http.StatusOK, projects)
}

func (con *ProjectController) GetAllProjectsIncludeHidden(context *gin.Context) {
	projects := con.projectService.GetAllProjects(true)
	if projects == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No projects found"})
		return
	}

	context.IndentedJSON(http.StatusOK, projects)
}

func (con *ProjectController) InsertProject(context *gin.Context) {
	var project *projectModel.ProjectDisplay
	if err := context.ShouldBindJSON(&project); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	createdProject, err := con.projectService.Insert(*project)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, createdProject)
}

func (con *ProjectController) UpdateProject(context *gin.Context) {
	var project *projectModel.ProjectDisplay

	if err := context.ShouldBindJSON(&project); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	updatedProject, err := con.projectService.Update(*project)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, updatedProject)
}

func (con *ProjectController) DeleteProject(context *gin.Context) {
	projectID, err := strconv.Atoi(context.Param("ID"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := con.projectService.Delete(projectID); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
