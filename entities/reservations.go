package entities

import "time"

type Reservation struct {
	BaseEntity
	UserID           string    `gorm:"size:36;index;not null" json:"user_id"`
	EquipmentID      string    `gorm:"size:36;index;not null" json:"equipment_id"`
	ResponsibleID    string    `gorm:"size:36;index" json:"responsible_id,omitempty"`
	ReservationStart time.Time `gorm:"not null" json:"reservation_start"`
	ReservationEnd   time.Time `gorm:"not null" json:"reservation_end"`

	// Relações (para queries com joins)
	// User        *User    `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	// Equipment     *Equipment `gorm:"foreignKey:equipmentID;references:ID" json:"equipment,omitempty"`
	// Responsible *User    `gorm:"foreignKey:ResponsibleID;references:ID" json:"responsible,omitempty"`
}
