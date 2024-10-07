package authRouter

import (
	authController "portfolioAPI/authentication/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(router *gin.Engine){
  authController := authController.NewAuthController()

  routerGroup := router.Group("/auth")
  {
    routerGroup.POST("signIn", authController.AuthenticateUser)
  }

  var existingAdmin = authController.IsAdminCreated()

  if !existingAdmin{
    routerGroup.POST("register", authController.RegisterAdmin)
  }
}
