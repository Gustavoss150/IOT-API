package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func ReservationRouter(router *gin.Engine) {
	r := router.Group("/api/reservations")
	{
		r.POST("/", controllers.CreateReservation)
		r.GET("/pending", controllers.ListPendingReservations)
		r.GET("/user/:userID", controllers.ListUserReservations)
		r.GET("/active", controllers.ListActiveReservations)
		r.GET("/day", controllers.ListApprovedReservationsByDay) // ex: GET /api/reservations/day?date=2024-01-15
		r.GET("/:id", controllers.GetReservation)
		r.PUT("/:id/approve", controllers.ApproveReservation)
		r.PUT("/:id/reject", controllers.RejectReservation)
	}
}
