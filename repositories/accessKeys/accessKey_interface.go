package accessKeysRepo

import "api/entities"

type AccessKeysRepository interface {
	Save(accessKey *entities.AccessKey) error
	Get(accessKey string) (*entities.AccessKey, error)
}
