package dto

import "time"

type ReservationDTO struct {
	UserID           string    `json:"user_id"`
	EquipmentID      string    `json:"equipment_id"`
	ResponsibleID    string    `json:"responsible_id"`
	ReservationStart time.Time `json:"reservation_start"`
	ReservationEnd   time.Time `json:"reservation_end"`
}
