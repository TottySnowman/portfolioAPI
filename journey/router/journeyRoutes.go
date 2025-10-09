package journeyRouter

import (
	"github.com/gin-gonic/gin"
	journeyController "portfolioAPI/journey/controllers"
	middleware "portfolioAPI/router/authorization"
)

func RegisterJourneyRoutes(router *gin.Engine, journeyController *journeyController.JourneyController) {
	routerGroup := router.Group("/journey")
	{
		routerGroup.GET("", journeyController.GetFullJourney)
		routerGroup.POST("", middleware.JWTMiddleware(), journeyController.InsertJourney)
    routerGroup.DELETE("/:ID", middleware.JWTMiddleware(), journeyController.DeleteExperience)
    routerGroup.PUT("/:ID", middleware.JWTMiddleware(), journeyController.UpdateExperience)
	}

}
