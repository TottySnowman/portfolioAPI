package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routerConfig "portfolioAPI/router"
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
