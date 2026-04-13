package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func AddSubjects(c *gin.Context) {
	var subjects models.Subjects
	if err := c.ShouldBindJSON(&subjects); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := services.AddSubjects(subjects.SubjectName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "subjects created"})
}
