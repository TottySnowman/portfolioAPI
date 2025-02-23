package statusRoutes

import (
	statusController "portfolioAPI/status/controllers"
	middleware "portfolioAPI/router/authorization"
	"github.com/gin-gonic/gin"
)

func RegisterTagRoutes(router *gin.Engine, statusController *statusController.StatusController) {
	routerGroup := router.Group("/status")
	{
		routerGroup.GET("",middleware.JWTMiddleware(), statusController.GetAllStatus)
	}
}

