package accessKeysRepo

import "api/entities"

type AccessKeysRepository interface {
	Save(accessKey *entities.AccessKey) error
}
