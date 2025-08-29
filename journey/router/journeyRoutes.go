package journeyRouter

import (
	journeyController "portfolioAPI/journey/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterJourneyRoutes(router *gin.Engine, journeyController *journeyController.JourneyController) {
	routerGroup := router.Group("/journey")
	{
		routerGroup.GET("", journeyController.GetFullJourney)
	}

}
