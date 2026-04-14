package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/controllers"
	"github.com/jarusuraj/schoolsystem/middlewares"
	"github.com/jarusuraj/schoolsystem/models"
)

func Profile(r *gin.Engine) {
	r.GET("/profile", middlewares.RateLimit(
		models.RateLimitConfig{
			Limit:  1,
			Status: "Request Failed",
			Body:   "API is at capacity, try again later.",
		}), func(c *gin.Context) {
		email, _ := c.Get("email")
		user_id, _ := c.Get("user_id")
		role, _ := c.Get("role")
		c.JSON(200, gin.H{"message": "this is the end hold your breath and count 10....", "email": email, "user_id": user_id, "role": role})
	})
}
func AddApprovalToUser(r *gin.Engine) {
	r.Use(middlewares.IsAdmin())
	r.PATCH("/admin/UserApproval", controllers.AddApproval)
}
