package server

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	routes.UserRouter(router)
}
