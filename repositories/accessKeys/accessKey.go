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
