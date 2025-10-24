package controllers

import (
	"api/dto"
	reservationRepo "api/repositories/reservations"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
	var reservationDTO dto.ReservationDTO
	if err := c.ShouldBindJSON(&reservationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório de reservas: " + err.Error()})
		return
	}

	reservation, err := services.CreateReservation(reservationDTO, repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Reserva criada com sucesso",
		"reservation": reservation,
	})
}
