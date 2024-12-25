package statusController

import (
	"net/http"
	statusService "portfolioAPI/status/services"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	statusService *statusService.StatusService
}

func NewStatusController(statusService *statusService.StatusService) *StatusController {
	return &StatusController{
		statusService: statusService,
	}
}

func (con *StatusController) GetAllStatus(context *gin.Context){
  
	existingStatus := con.statusService.GetAllStatus()

	if existingStatus == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No tags found"})
		return
	}

	context.IndentedJSON(http.StatusOK, existingStatus)

}
