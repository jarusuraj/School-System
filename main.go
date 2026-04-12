package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Problem in loading .env file.")
	}
	router := gin.Default()
	routes.Router(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Connection is OK."})
	})
	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal("Error while starting the server.")
	}
}
