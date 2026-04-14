package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func CreateClasses(c *gin.Context) {
	var req models.BulkCreateClassRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateClasses(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "classes created successfully",
	})
}