package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/blog-cms/controllers"
	"github.com/voduybaokhanh/blog-cms/middleware"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api/v1")

	// Auth
	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)

	// Protected
	api.GET("/me", middleware.AuthMiddleware(), controllers.Me)
}
