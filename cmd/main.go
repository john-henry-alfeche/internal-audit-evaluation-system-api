package main

import (
	"github.com/gin-gonic/gin"
	"internal-audit-evaluation-system-api/routes"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server on specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
