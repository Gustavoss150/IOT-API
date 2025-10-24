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

	usersRepo, err := usersRepo.InitUsersDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar banco de usuários"})
		return
	}

	createdUser, err := services.CreateUser(usersRepo, userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserResponse{
		ID:        createdUser.ID,
		DiscordID: createdUser.DiscordID,
		Email:     createdUser.Email,
		FullName:  createdUser.FullName,
		CreatedAt: createdUser.CreatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    response,
	})
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Necessário ID do usuário"})
		return
	}

	usersRepo, err := usersRepo.InitUsersDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar base de dados de usuários: " + err.Error()})
		return
	}

	user, err := services.GetUserProfile(usersRepo, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Erro ao buscar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuário": user})
}

func ListUsers(c *gin.Context) {
	usersRepo, err := usersRepo.InitUsersDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar base de dados de usuários: " + err.Error()})
		return
	}

	users, err := services.GetAllUsers(usersRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuários": users, "total": len(users)})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Necessário ID do usuário"})
		return
	}

	var updateDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida: " + err.Error()})
		return
	}

	usersRepo, err := usersRepo.InitUsersDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar base de dados de usuários"})
		return
	}

	if err := services.UpdateUser(usersRepo, updateDTO, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao atualizar usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuário atualizado": updateDTO})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário é obrigatório"})
		return
	}

	usersRepo, err := usersRepo.InitUsersDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar base de dados de usuários: " + err.Error()})
		return
	}

	if err := services.DeleteUser(usersRepo, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
