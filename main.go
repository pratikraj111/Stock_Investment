package main

import (
	"go-stock-app/config"
	"go-stock-app/server"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect database
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Set up routes
	router := gin.Default()
	server.SetupRoutes(router, config.DB)

	// Start server
	log.Fatal(router.Run(":" + config.ServerPort))
}
