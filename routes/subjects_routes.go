package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/controllers"
	"github.com/jarusuraj/schoolsystem/middlewares"
)

func AddSubjects(r *gin.Engine) {
	r.Use(middlewares.IsAdmin())
	r.POST("admin/addSubjects", controllers.AddSubjects)
}
