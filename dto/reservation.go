package dto

import "time"

type ReservationDTO struct {
	UserID           string    `json:"user_id" binding:"required"`
	EquipmentID      string    `json:"equipment_id" binding:"required"`
	ResponsibleID    string    `json:"responsible_id,omitempty"`
	ReservationStart time.Time `json:"reservation_start" binding:"required"`
	ReservationEnd   time.Time `json:"reservation_end" binding:"required"`
}
