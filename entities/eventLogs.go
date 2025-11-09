package entities

type ActionType string

const (
	Created ActionType = "created"
	Updated ActionType = "updated"
	Deleted ActionType = "deleted"
)

type EventLog struct {
	BaseEntity
	AccessKeyID   string     `gorm:"type:char(36);index" json:"access_key_id,omitempty"`
	UserID        string     `gorm:"type:char(36);index" json:"user_id,omitempty"`
	EquipmentID   string     `gorm:"type:char(36);index" json:"equipment_id,omitempty"`
	ReservationID string     `gorm:"type:char(36);index" json:"reservation_id,omitempty"`
	Action        ActionType `gorm:"type:ENUM('created','updated','deleted');not null" json:"action"`
	Message       string     `gorm:"type:text" json:"message,omitempty"`

	// Relações
	// AccessKey   *AccessKey   `gorm:"foreignKey:AccessKeyID;references:ID" json:"access_key,omitempty"`
	// User        *User        `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	// Equipment     *Equipment     `gorm:"foreignKey:EquipmentID;references:ID" json:"equipment,omitempty"`
	// Reservation *Reservation `gorm:"foreignKey:ReservationID;references:ID" json:"reservation,omitempty"`
}
