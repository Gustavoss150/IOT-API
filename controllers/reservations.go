package controllers

import (
	"api/dto"
	reservationRepo "api/repositories/reservations"
	"api/services"
	"net/http"
	"time"

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

func ListPendingReservations(c *gin.Context) {
	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório"})
		return
	}

	reservations, err := services.GetPendingReservations(repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservations": reservations,
		"total":        len(reservations),
	})
}

func ListUserReservations(c *gin.Context) {
	userID := c.Param("userID")

	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório"})
		return
	}

	reservations, err := services.GetUserReservations(repo, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservations": reservations,
		"total":        len(reservations),
	})
}

func ListActiveReservations(c *gin.Context) {
	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório de reservas: " + err.Error()})
		return
	}

	now := time.Now()

	reservations, err := services.GetActiveReservations(repo, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservations": reservations,
		"total":        len(reservations),
		"current_time": now,
	})
}

func ListApprovedReservationsByDay(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parâmetro 'date' é obrigatório (formato: YYYY-MM-DD)"})
		return
	}

	// Parse da data (formato: 2024-01-15)
	day, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "formato de data inválido. Use YYYY-MM-DD"})
		return
	}

	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório"})
		return
	}

	reservations, err := services.GetApprovedReservationsByDay(repo, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservations": reservations,
		"total":        len(reservations),
		"date":         day.Format("2006-01-02"),
	})
}

func GetReservation(c *gin.Context) {
	reservationID := c.Param("id")

	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório"})
		return
	}

	reservation, err := services.GetReservationByID(repo, reservationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservation": reservation})
}

func ApproveReservation(c *gin.Context) {
	reservationID := c.Param("id")
	responsibleID := c.GetHeader("admin-user-id")
	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório de reservas: " + err.Error()})
		return
	}
	err = services.ApproveReservation(repo, reservationID, responsibleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reserva aprovada com sucesso"})
}

func RejectReservation(c *gin.Context) {
	reservationID := c.Param("id")
	responsibleID := c.GetHeader("admin-user-id")
	repo, err := reservationRepo.InitReservationDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar repositório de reservas: " + err.Error()})
		return
	}
	err = services.RejectReservation(repo, reservationID, responsibleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reserva aprovada com sucesso"})
}
