package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func GetSubjects(c *gin.Context) {
	allSubs, err := services.GetSubjects()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to fetch subjects",
		})
		return
	}

	c.JSON(200, gin.H{
		"subjects": allSubs,
	})
}
func AddClassSubjects(c *gin.Context) {
	classNo, err := strconv.Atoi(c.Param("class_no"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid class_no"})
		return
	}

	var req models.BulkClassSubjectsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = services.AddClassSubjects(req, classNo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "subjects assigned successfully"})
}
