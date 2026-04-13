package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/middlewares"
	"github.com/jarusuraj/schoolsystem/routes"
)

func init() {
	gin.ForceConsoleColor()
}
func main() {
	config.LoadEnv()
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	defer config.DB.Close()
	router := gin.Default()
	router.Use(middlewares.CORS())
	routes.Router(router)
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("Error while starting the server.")
	}

}
