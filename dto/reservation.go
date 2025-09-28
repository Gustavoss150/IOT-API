package dto

import "time"

type ReservationDTO struct {
	UserID           string    `json:"user_id"`
	MachineID        string    `json:"machine_id"`
	ResponsibleID    string    `json:"responsible_id"`
	ReservationStart time.Time `json:"reservation_start"`
	ReservationEnd   time.Time `json:"reservation_end"`
}
