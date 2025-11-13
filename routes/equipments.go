package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func EquipmentRouter(router *gin.Engine) {
	r := router.Group("/api/equipments")
	{
		r.POST("/register", controllers.RegisterEquipment)
		r.GET("/:equipmentID", controllers.GetEquipmentByID)
		r.GET("/", controllers.ListEquipments)
		r.PUT("/:equipmentID", controllers.UpdateEquipment)
		r.DELETE("/:equipmentID", controllers.DeleteEquipment)
	}
}
