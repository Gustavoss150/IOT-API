package usersRepo

import "api/entities"

type UsersRepository interface {
	Save(user *entities.User) error
	GetByID(userID string) (*entities.User, error)
	GetByDiscordID(discordID string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	ExistsByDicord(discordID string) (bool, error)
	ExistsByEmail(email string) (bool, error)
	DeleteByID(id string) error
}
