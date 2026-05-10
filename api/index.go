package handler

import (
	"net/http"
	"sharing-vision-api/config"
	"sharing-vision-api/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	db := config.InitDB()
	// Kita tidak menutup DB di init karena ini serverless, 
	// koneksi akan di-manage oleh pool internal Go/MySQL.

	app = gin.Default()

	// CORS Middleware
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.SetupRoutes(app, db)
}

// Handler adalah entry point untuk Vercel Serverless Function
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
