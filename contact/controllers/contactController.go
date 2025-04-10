package contactController

import (
	"net/http"
	contactModels "portfolioAPI/contact/models"
	contactService "portfolioAPI/contact/services"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	service *contactService.ContactService
}

func NewContactController(service *contactService.ContactService) *ContactController {
	return &ContactController{
		service: service,
	}
}

func (con *ContactController) SendEmail(ctx *gin.Context) {
	var emailModel *contactModels.ContactModel
	if err := ctx.ShouldBindJSON(&emailModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	err := con.service.SendMail(*emailModel)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Failed to send email"})
		return
	}

	ctx.SetAccepted()
}
