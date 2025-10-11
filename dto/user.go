package dto

import "time"

type CreateUserDTO struct {
	DiscordID string `json:"discord_id" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	FullName  string `json:"full_name,omitempty"`
}

type UpdateUserDTO struct {
	DiscordID *string `json:"discord_id" binding:"required"`
	Email     *string `json:"email,omitempty" binding:"omitempty,email"`
	FullName  *string `json:"full_name,omitempty"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	DiscordID string    `json:"discord_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
