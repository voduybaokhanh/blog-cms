package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/controllers"
	"github.com/voduybaokhanh/blog-cms/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Auth
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.GET("/me", middleware.AuthMiddleware(), controllers.Me)

		// Users (chỉ admin mới có full CRUD)
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		{
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		// Posts
		posts := api.Group("/posts")
		posts.Use(middleware.AuthMiddleware())
		{
			posts.GET("", controllers.GetPosts)
			posts.GET("/:id", controllers.GetPost)
			posts.POST("", controllers.CreatePost)
			posts.PUT("/:id", controllers.UpdatePost) // chỉ admin hoặc chính author
			posts.DELETE("/:id", controllers.DeletePost)
		}

		// Categories
		categories := api.Group("/categories")
		categories.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		{
			categories.GET("", controllers.GetCategories)
			categories.POST("", controllers.CreateCategory)
			categories.DELETE(":id", controllers.DeleteCategory)
		}

		// Tags
		tags := api.Group("/tags")
		tags.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		{
			tags.GET("", controllers.GetTags)
			tags.POST("", controllers.CreateTag)
			tags.DELETE(":id", controllers.DeleteTag)
		}
	}

	return r
}
