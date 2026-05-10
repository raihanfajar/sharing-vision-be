package main

import (
	"log"
	"os"

	"sharing-vision-api/config"
	"sharing-vision-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	r := gin.Default()

	routes.SetupRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
