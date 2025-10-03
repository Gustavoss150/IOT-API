package entities

import (
	"time"
)

type BotConfig struct {
	BaseEntity
	OpeningTime                time.Time       `gorm:"type:time;not null" json:"opening_time"`
	ClosingTime                time.Time       `gorm:"type:time;not null" json:"closing_time"`
	MinReservation             int             `gorm:"not null;default:30" json:"min_reservation"`
	MaxReservationDays         int             `gorm:"not null;default:30" json:"max_reservation_days"`
	MaxReservationBlocks       int             `gorm:"not null;default:3" json:"max_reservation_blocks"`
	ReservationChannel         string          `gorm:"size:100" json:"reservation_channel,omitempty"`
	ReservationApprovalChannel string          `gorm:"size:100" json:"reservation_approval_channel,omitempty"`
	Holidays                   JSONStringSlice `gorm:"type:json" json:"holidays,omitempty"`
	DaysOfWeek                 JSONIntSlice    `gorm:"type:json" json:"days_of_week"`
	AdminRoles                 JSONStringSlice `gorm:"type:json" json:"admin_roles,omitempty"`
	ApproverRoles              JSONStringSlice `gorm:"type:json" json:"approver_roles,omitempty"`
}
