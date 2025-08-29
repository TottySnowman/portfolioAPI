package journeyController

import (
	"net/http"
	journeyService "portfolioAPI/journey/services"

	"github.com/gin-gonic/gin"
)

type JourneyController struct {
	journeyService *journeyService.JourneyService
}

func NewJourneyController(journeyService *journeyService.JourneyService) *JourneyController {
	return &JourneyController{
		journeyService: journeyService,
	}
}

func (con *JourneyController) GetFullJourney(context *gin.Context) {
	fullJourney := con.journeyService.GetFullJourney()
	context.IndentedJSON(http.StatusOK, fullJourney)
}
