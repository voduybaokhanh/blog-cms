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
	}

	return r
}
