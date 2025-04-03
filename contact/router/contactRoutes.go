package contactRouter

import (
	contactController "portfolioAPI/contact/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterContactRoutes(router *gin.Engine, contactController *contactController.ContactController) {
	routerGroup := router.Group("")
	{
		routerGroup.POST("/contact", contactController.SendEmail)
	}
}
