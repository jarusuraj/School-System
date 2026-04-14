package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/middlewares"
)

func Router(router *gin.Engine) {
	AuthRoutes(router)
	router.Use(middlewares.AuthGate())
	Profile(router)
	AddSubjects(router)
	AddApprovalToUser(router)
	Class(router)
}
