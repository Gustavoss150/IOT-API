package entities

type EquipmentStatus string

const (
	StatusAvailable   EquipmentStatus = "available"
	StatusInUse       EquipmentStatus = "in_use"
	StatusMaintenance EquipmentStatus = "maintenance"
)

type Equipment struct {
	BaseEntity
	Name        string          `gorm:"size:100;not null" json:"name"`
	Description string          `gorm:"type:text" json:"description,omitempty"`
	Status      EquipmentStatus `gorm:"type:ENUM('available', 'in_use', 'maintenance'); default:'available'" json:"status"`
}
