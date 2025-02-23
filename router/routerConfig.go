package router

import (
	"os"
	authRouter "portfolioAPI/authentication/router"
	chatRoutes "portfolioAPI/chat/router"
	dependencyinjection "portfolioAPI/dependencyInjection"
	fileRoutes "portfolioAPI/fileUpload/router"
	knowledgeRoutes "portfolioAPI/knowledge/router"
	project_routes "portfolioAPI/project/router"
	statusRoutes "portfolioAPI/status/router"
	tagRoutes "portfolioAPI/tag/router"

	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, appContainer *dependencyinjection.AppContainer) {
	setupCors(router)
	setupRoutes(router, appContainer)
	setupLogoServ(router)
	setupCvServ(router)
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

func setupRoutes(router *gin.Engine, appContainer *dependencyinjection.AppContainer) {
	project_routes.RegisterProjectRoutes(router, appContainer.ProjectController)
	tagRoutes.RegisterTagRoutes(router, appContainer.TagController)
	statusRoutes.RegisterTagRoutes(router, appContainer.StatusController)
	chatRoutes.RegisterChatRoutes(router, appContainer.ChatController)
	fileRoutes.RegisterFileRoutes(router, appContainer.FileController)
	knowledgeRoutes.RegisterKnowledgeRoutes(router, appContainer.KnowledgeController)
	authRouter.RegisterAuthRouter(router)
}

func setupLogoServ(router *gin.Engine) {
	router.Static("/logo", "./public/logo")
}

func setupCvServ(router *gin.Engine) {
	router.Static("/cv", "./public/cv")
}
