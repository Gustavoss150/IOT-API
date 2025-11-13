package services

import (
	"api/dto"
	"api/entities"
	equipmentsRepo "api/repositories/equipments"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateEquipment(equipmentsRepo equipmentsRepo.EquipmentsRepository, equipmentDTO dto.CreateEquipmentDTO) (*entities.Equipment, error) {
	if equipmentDTO.Name == "" {
		return nil, errors.New("o nome do equipamento é obrigatório")
	}

	// Validação do status
	status := entities.EquipmentStatus(equipmentDTO.Status)
	if status == "" {
		status = entities.StatusAvailable // valor padrão
	} else if status != entities.StatusAvailable && status != entities.StatusInUse && status != entities.StatusMaintenance {
		return nil, fmt.Errorf("status inválido: %s", equipmentDTO.Status)
	}

	equipment := &entities.Equipment{
		Name:        equipmentDTO.Name,
		Description: equipmentDTO.Description,
		Status:      status,
	}

	if err := equipmentsRepo.Save(equipment); err != nil {
		return nil, fmt.Errorf("erro ao salvar equipamento: %v", err)
	}

	return equipment, nil
}

func GetEquipmentProfile(equipmentsRepo equipmentsRepo.EquipmentsRepository, equipmentID string) (*entities.Equipment, error) {
	equipment, err := equipmentsRepo.GetByID(equipmentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("equipamento não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar equipamento: %v", err)
	}
	return equipment, nil
}

func GetAllEquipments(equipmentsRepo equipmentsRepo.EquipmentsRepository) ([]*entities.Equipment, error) {
	equipments, err := equipmentsRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar equipamentos: %v", err)
	}

	if len(equipments) == 0 {
		return nil, errors.New("nenhum equipamento encontrado")
	}

	return equipments, nil
}

func UpdateEquipment(equipmentsRepo equipmentsRepo.EquipmentsRepository, updateRequest dto.UpdateEquipmentDTO, targetEquipmentID string) error {
	equipment, err := equipmentsRepo.GetByID(targetEquipmentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("equipamento não encontrado")
		}
		return fmt.Errorf("erro ao buscar equipamento: %v", err)
	}

	if updateRequest.Name != nil {
		equipment.Name = *updateRequest.Name
	}

	if updateRequest.Description != nil {
		equipment.Description = *updateRequest.Description
	}

	if updateRequest.Status != nil {
		status := entities.EquipmentStatus(*updateRequest.Status)
		if status != entities.StatusAvailable && status != entities.StatusInUse && status != entities.StatusMaintenance {
			return fmt.Errorf("status inválido: %s", *updateRequest.Status)
		}
		equipment.Status = status
	}

	if err := equipmentsRepo.Save(equipment); err != nil {
		return fmt.Errorf("erro ao atualizar equipamento: %v", err)
	}

	return nil
}

func DeleteEquipment(equipmentsRepo equipmentsRepo.EquipmentsRepository, equipmentID string) error {
	if equipmentID == "" {
		return errors.New("id é obrigatório")
	}

	err := equipmentsRepo.DeleteByID(equipmentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("equipamento não encontrado")
		}
		return fmt.Errorf("erro ao deletar equipamento: %v", err)
	}

	return nil
}
