package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	r := router.Group("/api/users")
	{
		r.POST("/register", controllers.RegisterUser)
	}
}
