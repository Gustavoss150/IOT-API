package entities

import "time"

type StatusReservation string

const (
	Approved StatusReservation = "approved"
	Rejected StatusReservation = "rejected"
	Pending  StatusReservation = "pending"
)

type Reservation struct {
	BaseEntity
	UserID           string            `gorm:"type:char(36);index;not null" json:"user_id"`
	EquipmentID      string            `gorm:"type:char(36);index;not null" json:"equipment_id"`
	ResponsibleID    string            `gorm:"type:char(36);index" json:"responsible_id,omitempty"`
	ReservationStart time.Time         `gorm:"not null" json:"reservation_start"`
	ReservationEnd   time.Time         `gorm:"not null" json:"reservation_end"`
	Status           StatusReservation `gorm:"type:ENUM('approved','rejected','pending'); default:'pending'" json:"status"`

	// Relações (para queries com joins)
	// User        *User    `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	// Equipment     *Equipment `gorm:"foreignKey:equipmentID;references:ID" json:"equipment,omitempty"`
	// Responsible *User    `gorm:"foreignKey:ResponsibleID;references:ID" json:"responsible,omitempty"`
}
