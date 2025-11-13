package controllers

import (
	"api/dto"
	equipmentsRepo "api/repositories/equipments"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEquipment(c *gin.Context) {
	var equipmentDTO dto.CreateEquipmentDTO
	if err := c.ShouldBindJSON(&equipmentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	repo, err := equipmentsRepo.InitEquipmentsDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de dados de equipamentos: " + err.Error()})
		return
	}

	createdEquipment, err := services.CreateEquipment(repo, equipmentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dto.EquipmentResponse{
		ID:          createdEquipment.ID,
		Name:        createdEquipment.Name,
		Description: createdEquipment.Description,
		Status:      string(createdEquipment.Status),
		CreatedAt:   createdEquipment.CreatedAt,
		UpdatedAt:   createdEquipment.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Equipamento criado com sucesso",
		"equipment": response,
	})
}

func GetEquipmentByID(c *gin.Context) {
	equipmentID := c.Param("equipmentID")
	if equipmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do equipamento é obrigatório"})
		return
	}

	repo, err := equipmentsRepo.InitEquipmentsDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de dados: " + err.Error()})
		return
	}

	equipment, err := services.GetEquipmentProfile(repo, equipmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.EquipmentResponse{
		ID:          equipment.ID,
		Name:        equipment.Name,
		Description: equipment.Description,
		Status:      string(equipment.Status),
		CreatedAt:   equipment.CreatedAt,
		UpdatedAt:   equipment.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"equipment": response})
}

func ListEquipments(c *gin.Context) {
	repo, err := equipmentsRepo.InitEquipmentsDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de dados: " + err.Error()})
		return
	}

	equipments, err := services.GetAllEquipments(repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":      len(equipments),
		"equipments": equipments,
	})
}

func UpdateEquipment(c *gin.Context) {
	equipmentID := c.Param("equipmentID")
	if equipmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do equipamento é obrigatório"})
		return
	}

	var updateDTO dto.UpdateEquipmentDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	repo, err := equipmentsRepo.InitEquipmentsDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de dados: " + err.Error()})
		return
	}

	if err := services.UpdateEquipment(repo, updateDTO, equipmentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipamento atualizado com sucesso"})
}

func DeleteEquipment(c *gin.Context) {
	equipmentID := c.Param("equipmentID")
	if equipmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do equipamento é obrigatório"})
		return
	}

	repo, err := equipmentsRepo.InitEquipmentsDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de dados: " + err.Error()})
		return
	}

	if err := services.DeleteEquipment(repo, equipmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
