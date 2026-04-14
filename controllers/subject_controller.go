package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func AddSubjects(c *gin.Context) {
	var req models.BulkCreateSubjectsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := services.AddSubjects(req.Subjects)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "subjects processed successfully",
		"subjects": result,
	})
}
