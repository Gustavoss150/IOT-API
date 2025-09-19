package dto

import "time"

type ReservationDTO struct {
	MachineID        string    `json:"machine_id"`
	ResponsibleID    string    `json:"responsible_id"`
	ReservationStart time.Time `json:"reservation_start"`
	ReservationEnd   time.Time `json:"reservation_end"`
}
