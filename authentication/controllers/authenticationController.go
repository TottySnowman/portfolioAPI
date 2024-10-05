package authController

import (
	"log"
	"net/http"
	authenticationModel "portfolioAPI/authentication/models"
	authService "portfolioAPI/authentication/service"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	service *authService.AuthService
}

func NewAuthController() *AuthenticationController {
	return &AuthenticationController{
		service: authService.NewAuthService(),
	}
}

func (con *AuthenticationController) AuthenticateUser(context *gin.Context) {
	var userInput authenticationModel.LoginRequest
	result := con.service.AuthenticateUser(userInput)

	if result == nil{
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

  context.IndentedJSON(http.StatusOK, gin.H{"token": result.Token})
}
