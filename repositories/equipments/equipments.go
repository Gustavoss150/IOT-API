package equipmentsRepo

import (
	"api/config"
	"api/entities"
	"errors"

	"gorm.io/gorm"
)

type equipmentsRepository struct {
	DB *gorm.DB
}

func InitEquipmentsDatabase() (EquipmentsRepository, error) {
	db := config.DB
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}
	return &equipmentsRepository{DB: db}, nil
}

func (r *equipmentsRepository) Save(equipment *entities.Equipment) error {
	return r.DB.Save(equipment).Error
}

func (r *equipmentsRepository) GetByID(equipmentID string) (*entities.Equipment, error) {
	var equipment entities.Equipment
	if err := r.DB.Where("id = ?", equipmentID).First(&equipment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("equipamento não encontrado")
		}
		return nil, err
	}
	return &equipment, nil
}

func (r *equipmentsRepository) GetAll() ([]*entities.Equipment, error) {
	var equipments []*entities.Equipment
	if err := r.DB.Order("created_at DESC").Find(&equipments).Error; err != nil {
		return nil, err
	}
	return equipments, nil
}

func (r *equipmentsRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := r.DB.Model(&entities.Equipment{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *equipmentsRepository) DeleteByID(id string) error {
	result := r.DB.Delete(&entities.Equipment{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
