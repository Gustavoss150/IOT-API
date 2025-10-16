package main

import (
	"api/config"
	"api/migrations"
	"api/server"
	"api/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	migrations.Migrate()

	r := gin.Default()

	server.InitRoutes(r)

	go services.StartMQTT()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is working"})
	})
	r.Run(":8080")
}
