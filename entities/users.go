package entities

type User struct {
	BaseEntity
	DiscordID string `gorm:"size:255;uniqueIndex" json:"discord_id,omitempty"`
	FullName  string `gorm:"size:255" json:"full_name,omitempty"`
	Email     string `gorm:"size:150;uniqueIndex;not null" json:"email"`
}
