package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.Get("email")
		role, _ := c.Get("role")
		if email != os.Getenv("ADMIN_EMAIL") && role != os.Getenv("ADMIN_ROLE") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			c.Abort()
		}
		c.Next()
	}
}
