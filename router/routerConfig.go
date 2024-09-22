package router

import (
	"os"
	project_routes "portfolioAPI/project/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	setupCors(router)
	setupRoutes(router)
	setupLogoServ(router)
}

func setupCors(router *gin.Engine) {
	corsOrigin := os.Getenv("CORS_ORIGIN1")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

}

func setupRoutes(router *gin.Engine) {
	project_routes.RegisterProjectRoutes(router)
}

func setupLogoServ(router *gin.Engine) {
	router.Static("/logo", "./logo")
}
