package controllers

import (
	"api/dto"
	usersRepo "api/repositories/users"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	usersRepo, err := usersRepo.InitUsersRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de usuários"})
		return
	}

	createdUser, err := services.CreateUser(usersRepo, userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    createdUser,
	})
}
