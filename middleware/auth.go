package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Normalize user_id to uint in context
		var userIDUint uint
		switch v := claims["user_id"].(type) {
		case float64:
			userIDUint = uint(v)
		case int:
			userIDUint = uint(v)
		case int64:
			userIDUint = uint(v)
		case uint:
			userIDUint = v
		case uint64:
			userIDUint = uint(v)
		case string:
			var n uint
			for i := 0; i < len(v); i++ {
				ch := v[i]
				if ch < '0' || ch > '9' {
					n = 0
					break
				}
				n = n*10 + uint(ch-'0')
			}
			userIDUint = n
		default:
			userIDUint = 0
		}

		if userIDUint == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token user"})
			c.Abort()
			return
		}

		c.Set("user_id", userIDUint)
		c.Set("role", claims["role"])
		c.Next()
	}
}
