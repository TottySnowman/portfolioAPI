package fileRoutes

import (
	fileController "portfolioAPI/fileUpload/controllers"
	middleware "portfolioAPI/router/authorization"

	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(router *gin.Engine, fileController *fileController.FileController) {
	routerGroup := router.Group("")
	{
		routerGroup.POST("/logo", middleware.JWTMiddleware(), fileController.UploadFile)
		routerGroup.POST("/cv", middleware.JWTMiddleware(), fileController.UploadFile)
		routerGroup.DELETE("/logo", middleware.JWTMiddleware(), fileController.UploadFile)
		routerGroup.DELETE("/cv", middleware.JWTMiddleware(), fileController.UploadFile)
	}
}
