package usersRepo

import (
	"api/config"
	"api/entities"
	"errors"

	"gorm.io/gorm"
)

type usersRepository struct {
	DB *gorm.DB
}

func InitUsersDatabase() (UsersRepository, error) {
	db := config.DB
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}
	return &usersRepository{DB: db}, nil
}

func (r *usersRepository) Save(user *entities.User) error {
	return r.DB.Save(user).Error
}

func (r *usersRepository) GetByID(userID string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}

func (r *usersRepository) GetByDiscordID(discordID string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.Where("discord_id = ?", discordID).First(&user).Error; err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}

func (r *usersRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}

func (r *usersRepository) GetAll() ([]*entities.User, error) {
	var users []*entities.User
	if err := r.DB.Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *usersRepository) ExistsByDicord(discordID string) (bool, error) {
	var count int64
	err := r.DB.Model(&entities.User{}).Where("discord_id = ?", discordID).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *usersRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&entities.User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *usersRepository) DeleteByID(id string) error {
	result := r.DB.Delete(&entities.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
