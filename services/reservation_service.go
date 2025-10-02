package services

import (
	"api/dto"
	"api/entities"
	reservationRepo "api/repositories/reservations"
	"errors"
)

func CreateReservation(reservationDTO dto.ReservationDTO, repo reservationRepo.ReservationRepository) (*entities.Reservation, error) {
	if reservationDTO.UserID == "" || reservationDTO.EquipmentID == "" {
		return nil, errors.New("id do usuário e da máquina não podem ser nulos")
	}

	if reservationDTO.ReservationEnd.Before(reservationDTO.ReservationStart) || reservationDTO.ReservationEnd.Equal(reservationDTO.ReservationStart) {
		return nil, errors.New("horário de término da reserva deve ser após o horário de início")
	}

	conflict, err := repo.HasReservationConflict(reservationDTO.EquipmentID, reservationDTO.ReservationStart, reservationDTO.ReservationEnd)
	if err != nil {
		return nil, err
	}
	if conflict {
		return nil, errors.New("já existe reserva para esse horário")
	}

	reservation := &entities.Reservation{
		UserID:           reservationDTO.UserID,
		EquipmentID:      reservationDTO.EquipmentID,
		ResponsibleID:    reservationDTO.ResponsibleID,
		ReservationStart: reservationDTO.ReservationStart,
		ReservationEnd:   reservationDTO.ReservationEnd,
	}

	if err := repo.Save(reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

// dto será alimentado e chamado pelo controller
