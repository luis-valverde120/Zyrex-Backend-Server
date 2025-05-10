package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Middleware global
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func StartServer(router *gin.Engine, port string) {
	// Start the server on port 8080
	fmt.Print("Starting server on port " + port + "...\n")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
