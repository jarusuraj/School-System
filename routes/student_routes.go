package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/middlewares"
	"github.com/jarusuraj/schoolsystem/models"
)

func Profile(r *gin.Engine) {
	r.GET("/profile", middlewares.RateLimit(
		models.RateLimitConfig{
			Limit:  60.0 / 60.0,
			Status: "Request Failed",
			Body:   "API is at capacity, try again later.",
		}), func(c *gin.Context) {
		email, _ := c.Get("email")
		user_id, _ := c.Get("user_id")
		role, _ := c.Get("role")
		c.JSON(200, gin.H{"message": "this is the end hold your breath and count 10....", "email": email, "user_id": user_id, "role": role})
	})
}
