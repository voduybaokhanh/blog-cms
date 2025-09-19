package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/voduybaokhanh/blog-cms/config"
	"github.com/voduybaokhanh/blog-cms/controllers"
	"github.com/voduybaokhanh/blog-cms/models"
	"github.com/voduybaokhanh/blog-cms/routes"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	// Kết nối DB
	config.ConnectDatabase()

	// Auto migrate
	err := config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}, &models.Tag{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	// Gắn DB vào controllers
	controllers.SetDB(config.DB) // cho auth.go

	// Setup router (dùng routes.go, có cả /users)
	r := routes.SetupRouter()

	// Run
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
