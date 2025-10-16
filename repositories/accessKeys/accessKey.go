package accessKeysRepo

import (
	"api/config"
	"api/entities"
	"errors"

	"gorm.io/gorm"
)

type accessKeysRepository struct {
	DB *gorm.DB
}

func InitAccessKeyDatabase() (AccessKeysRepository, error) {
	db := config.DB
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}
	return &accessKeysRepository{DB: db}, nil
}

func (r *accessKeysRepository) Save(accessKey *entities.AccessKey) error {
	return r.DB.Save(accessKey).Error
}

func (r *accessKeysRepository) Get(accessKey string) (*entities.AccessKey, error) {
	var key entities.AccessKey
	if err := r.DB.Where("id = ?", accessKey).First(&key).Error; err != nil {
		return nil, errors.New("chave não encontrada")
	}
	return &key, nil
}
