package entities

type TypeKey string
type StatusKey string

const (
	Closed StatusKey = "closed"
	Open   StatusKey = "open"
)

const (
	Rfid   TypeKey = "rfid"
	Qrcode TypeKey = "qrcode"
	Pin    TypeKey = "pin"
	Other  TypeKey = "other"
)

type AccessKey struct {
	BaseEntity
	TypeKey       TypeKey   `gorm:"type:ENUM('rfid','qrcode','pin','other');not null" json:"type_key"`
	StatusKey     StatusKey `gorm:"type:ENUM('closed','open');not null" json:"status_key"`
	Value         string    `gorm:"size:255;uniqueIndex;notnull" json:"value"`
	AssignedTo    string    `gorm:"size:36;index" json:"assigned_to,omitempty"`
	ReservationID string    `gorm:"size:36;index" json:"reservation_id,omitempty"`
	// IsActive (desativado por padrão) para ativar chave quando iniciar o horário de reserva
	// IsActive		bool	`json:"is_active"`

	// Relações
	// User        *User        `gorm:"foreignKey:AssignedTo;references:ID" json:"user,omitempty"`
	// Reservation *Reservation `gorm:"foreignKey:ReservationID;references:ID" json:"reservation,omitempty"`
}
