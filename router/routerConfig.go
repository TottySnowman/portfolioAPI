package router

import (
	"os"
	authRouter "portfolioAPI/authentication/router"
	project_routes "portfolioAPI/project/router"
	tagRoutes "portfolioAPI/tag/router"

	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	setupCors(router)
	setupRoutes(router)
	setupLogoServ(router)
}

func setupCors(router *gin.Engine) {
	corsOrigin0 := os.Getenv("CORS_ORIGIN0")
	corsOrigin1 := os.Getenv("CORS_ORIGIN1")

	release := os.Getenv("GIN_MODE")
	if strings.Compare(release, "release") == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{corsOrigin0, corsOrigin1},
		AllowMethods:     []string{"GET", "DELETE", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

}

func setupRoutes(router *gin.Engine) {
	project_routes.RegisterProjectRoutes(router)
  tagRoutes.RegisterTagRoutes(router)

	authRouter.RegisterAuthRouter(router)
}

func setupLogoServ(router *gin.Engine) {
	router.Static("/logo", "./logo")
}
