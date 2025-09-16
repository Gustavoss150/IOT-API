package entities

import "time"

type Reservation struct {
	BaseEntity
	UserID           string    `gorm:"size:36;index;not null" json:"user_id"`
	MachineID        string    `gorm:"size:36;index;not null" json:"machine_id"`
	ResponsibleID    string    `gorm:"size:36;index" json:"responsible_id,omitempty"`
	ReservationStart time.Time `gorm:"not null" json:"reservation_start"`
	ReservationEnd   time.Time `gorm:"not null" json:"reservation_end"`

	// Relações (para queries com joins)
	// User        *User    `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	// Machine     *Machine `gorm:"foreignKey:machineID;references:ID" json:"machine,omitempty"`
	// Responsible *User    `gorm:"foreignKey:ResponsibleID;references:ID" json:"responsible,omitempty"`
}
