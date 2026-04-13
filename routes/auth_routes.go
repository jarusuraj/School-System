package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/controllers"
	"github.com/jarusuraj/schoolsystem/middlewares"
	"github.com/jarusuraj/schoolsystem/models"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/signup", middlewares.RateLimit(
		models.RateLimitConfig{
			Limit:  58.0 / 60.0,
			Status: "Request Failed",
			Body:   "API is at capacity, try again later.",
		}), controllers.Signup)

	auth.POST("/login", middlewares.RateLimit(
		models.RateLimitConfig{
			Limit:  5.0 / 60.0,
			Status: "Request Failed",
			Body:   "API is at capacity, try again later.",
		}), controllers.Login)
}
