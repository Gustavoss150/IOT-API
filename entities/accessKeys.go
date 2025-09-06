package entities

type TypeKey string

const (
	Rfid   TypeKey = "rfid"
	Qrcode TypeKey = "qrcode"
	Pin    TypeKey = "pin"
	Other  TypeKey = "other"
)

type AccessKey struct {
	BaseEntity
	TypeKey       TypeKey `gorm:"type:ENUM('rfid','qrcode','pin','other');not null" json:"type_key"`
	Value         string  `gorm:"size:255;uniqueIndex;notnull" json:"value"`
	AssignedTo    string  `gorm:"size:36;index" json:"assigned_to,omitempty"`
	ReservationID string  `gorm:"size:36;index" json:"reservation_id,omitempty"`

	// Relações
	User        User        `gorm:"foreignKey:AssignedTo" json:"user,omitempty"`
	Reservation Reservation `gorm:"foreignKey:ReservationID" json:"reservation,omitempty"`
}
