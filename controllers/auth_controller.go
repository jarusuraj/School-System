package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func Signup(c *gin.Context) {
	var user models.SignupRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := services.Signup(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Signup successfull."})

}
func Login(c *gin.Context) {
	var login models.LoginRequest
	c.ShouldBindJSON(&login)
	token, err := services.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(401, gin.H{"message": "login failed."})
	}
	c.JSON(200, gin.H{"message": "Login Successfull", "Token": token})
}
