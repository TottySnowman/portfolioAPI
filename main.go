package main

import (
	"log"
	"os"

	dependencyinjection "portfolioAPI/dependencyInjection"
	routerConfig "portfolioAPI/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env variables")
	}

	router := gin.Default()

  appContainer := dependencyinjection.NewAppContainer()
	routerConfig.SetupRouter(router, appContainer)
	serverPort := os.Getenv("SERVER_PORT")

	router.Run(":" + serverPort)
}
