package main

import (
	"log"
	"os"

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

	routerConfig.SetupRouter(router)
	serverPort := os.Getenv("SERVER_PORT")

	router.Run(":" + serverPort)
}
