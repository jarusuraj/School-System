package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	defer config.DB.Close()
	router := gin.Default()
	routes.Router(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Connection is OK."})
	})
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("Error while starting the server.")
	}

}
