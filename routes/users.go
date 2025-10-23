package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	r := router.Group("/api/users")
	{
		r.POST("/register", controllers.RegisterUser)
		r.GET("/:userID", controllers.GetUserByID)
		r.GET("/", controllers.ListUsers)
		r.PUT("/:userID", controllers.UpdateUser)
		r.DELETE("/:userID", controllers.DeleteUser)
	}
}
