package main

import (
	"log"
	"os"

	gin "github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Gin router
	router := gin.Default()

	// Start server
	log.Println("Server running on port:", port)
	router.Run(":" + port)
}
