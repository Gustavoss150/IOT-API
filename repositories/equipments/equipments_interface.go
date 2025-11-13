package equipmentsRepo

import "api/entities"

type EquipmentsRepository interface {
	Save(equipment *entities.Equipment) error
	GetByID(equipmentID string) (*entities.Equipment, error)
	GetAll() ([]*entities.Equipment, error)
	ExistsByName(name string) (bool, error)
	DeleteByID(id string) error
}
