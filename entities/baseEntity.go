package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// campos base para todas as entidades
type BaseEntity struct {
	ID        string     `json:"id" gorm:"primaryKey;type:char(36)"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}
