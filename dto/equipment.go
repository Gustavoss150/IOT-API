package dto

import "time"

type CreateEquipmentDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type UpdateEquipmentDTO struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type EquipmentResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
