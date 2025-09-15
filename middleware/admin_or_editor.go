package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOrEditor() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role != "admin" && role != "editor") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admins or Editors only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
