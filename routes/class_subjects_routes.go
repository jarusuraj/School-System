package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/controllers"
	"github.com/jarusuraj/schoolsystem/middlewares"
)

func Class(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middlewares.IsAdmin())

	admin.GET("/subjects", controllers.GetSubjects)
	admin.POST("/class/:class_no/subjects", controllers.AddClassSubjects)
	admin.POST("/classes", controllers.CreateClasses)
}
