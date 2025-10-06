package controllers

import (
	"api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterReservation(c *gin.Context) {
	var reservationDTO dto.ReservationDTO
	if err := c.ShouldBindJSON(&reservationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// resto ainda a fazer, necessário service de accessKey
}
