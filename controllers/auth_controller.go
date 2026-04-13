package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/services"
)

func Signup(c *gin.Context) {
	var user models.SignupRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := services.Signup(user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Signup successfull."})

}
func Login(c *gin.Context) {
	var login models.LoginRequest
	c.ShouldBindJSON(&login)
	token, err := services.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(401, gin.H{"message": "login failed. " + err.Error()})
		return
	}
	c.SetCookie(
		"token",
		token,
		3600,  // 1 hour
		"/",   //which path can use
		"",    //which domain can use here it is using the default localhost
		false, // true in production (HTTPS)
		true,  // HttpOnly (important)
	)
	c.JSON(200, gin.H{"message": "Login Successfull"})
}
