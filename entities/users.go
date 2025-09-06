package entities

type UserRole string

const (
	RoleStudent  UserRole = "student"
	RoleTeacher  UserRole = "teacher"
	RoleCustomer UserRole = "customer"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	BaseEntity
	FullName string   `gorm:"size:255;not null" json:"full_name"`
	Email    string   `gorm:"size:150;uniqueIndex;not null" json:"email"`
	Username string   `gorm:"size:150;uniqueIndex;not null" json:"username"`
	Password string   `gorm:"size:255" json:"password,omitempty"`
	Role     UserRole `gorm:"type:ENUM('student','teacher','customer','admin');not null" json:"role"`
}
