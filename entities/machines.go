package entities

type MachineStatus string

const (
	StatusAvailable   MachineStatus = "available"
	StatusInUse       MachineStatus = "in_use"
	StatusMaintenance MachineStatus = "maintenance"
)

type Machine struct {
	BaseEntity
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"type:text" json:"description,omitempty"`
	Status      MachineStatus `gorm:"type:ENUM('available', 'in_use', 'maintenance'); default:'available'" json:"status"`
}
