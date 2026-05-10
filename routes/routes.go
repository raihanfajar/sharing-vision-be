package routes

import (
	"database/sql"
	"time"

	"sharing-vision-api/controllers"
	"sharing-vision-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	postService := services.NewPostService(db)
	postController := controllers.NewPostController(postService)

	r.POST("/article/", postController.CreatePost)
	r.GET("/article/list/:limit/:offset", postController.GetAllPosts)
	r.GET("/article/:id", postController.GetPostByID)
	r.PUT("/article/:id", postController.UpdatePost)
	r.DELETE("/article/:id", postController.DeletePost)
}
