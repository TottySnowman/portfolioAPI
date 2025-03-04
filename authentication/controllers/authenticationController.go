package authController

import (
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
	var userInput *authenticationModel.LoginRequest
	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	result := con.service.AuthenticateUser(userInput)

	if result == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

  context.IndentedJSON(http.StatusOK, gin.H{"token": result.Token, "Username": result.Username})
}

func (con *AuthenticationController) IsAdminCreated() bool {
	return con.service.IsAdminCreated()
}

func (con *AuthenticationController) RegisterAdmin(context *gin.Context) {
  if con.service.IsAdminCreated(){
    context.JSON(http.StatusBadRequest, gin.H{ "message": "Admin already exists" })
    return
  }

  if con.service.RegisterAdmin(){
    context.Status(http.StatusCreated)
    return
  }

  context.JSON(http.StatusBadRequest, gin.H{ "message": "Failed to create admin!" })
}
