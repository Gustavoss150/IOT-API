package main

import (
	"api/config"
	"api/migrations"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	migrations.Migrate()

	r := gin.Default()

	// server.InitRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is working"})
	})
	r.Run(":8080")
}
