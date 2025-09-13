package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/config"
	"github.com/voduybaokhanh/blog-cms/controllers"
	"github.com/voduybaokhanh/blog-cms/models"
)

func main() {
	// Kết nối DB
	config.ConnectDatabase()

	// Auto migrate
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	// Gắn DB vào controllers
	controllers.SetDB(config.DB)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.GET("/me", controllers.Me)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
