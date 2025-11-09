package services

import (
	"api/dto"
	"api/entities"
	reservationRepo "api/repositories/reservations"
	"errors"
	"time"
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

func GetReservationByID(repo reservationRepo.ReservationRepository, reservationID string) (*entities.Reservation, error) {
	return repo.GetByID(reservationID)
}

func GetUserReservations(repo reservationRepo.ReservationRepository, userID string) ([]entities.Reservation, error) {
	if userID == "" {
		return nil, errors.New("ID do usuário é obrigatório")
	}
	return repo.GetByUserID(userID)
}

func GetActiveReservations(repo reservationRepo.ReservationRepository, now time.Time) ([]entities.Reservation, error) {
	allApproved, err := repo.GetReservationsByStatus(entities.Approved)
	if err != nil {
		return nil, err
	}

	var active []entities.Reservation
	for _, r := range allApproved {
		if IsReservationActive(r, now) {
			active = append(active, r)
		}
	}
	return active, nil
}

func GetApprovedReservationsByDay(repo reservationRepo.ReservationRepository, day time.Time) ([]entities.Reservation, error) {
	if day.IsZero() {
		return nil, errors.New("data é obrigatória")
	}
	return repo.GetApprovedReservationsByDay(day)
}

func GetPendingReservations(repo reservationRepo.ReservationRepository) ([]entities.Reservation, error) {
	return repo.GetReservationsByStatus(entities.Pending)
}

func ApproveReservation(repo reservationRepo.ReservationRepository, reservationID string, responsibleID string) error {
	reservation, err := repo.GetByID(reservationID)
	if err != nil {
		return errors.New("reserva não encontrada")
	}

	if reservation.Status != entities.Pending {
		return errors.New("apenas reservas pendentes podem ser aprovadas")
	}

	conflict, err := repo.HasReservationConflict(reservation.EquipmentID, reservation.ReservationStart, reservation.ReservationEnd)
	if err != nil {
		return err
	}
	if conflict {
		return errors.New("não é possível aprovar - conflito com reserva aprovada existente")
	}

	reservation.Status = entities.Approved
	reservation.ResponsibleID = responsibleID
	return repo.Save(reservation)
}

func RejectReservation(repo reservationRepo.ReservationRepository, reservationID string, responsibleID string) error {
	reservation, err := repo.GetByID(reservationID)
	if err != nil {
		return errors.New("reserva não encontrada")
	}

	if reservation.Status != entities.Pending {
		return errors.New("apenas reservas pendentes podem ser rejeitadas")
	}

	reservation.Status = entities.Rejected
	reservation.ResponsibleID = responsibleID
	return repo.Save(reservation)
}

func IsReservationActive(reservation entities.Reservation, now time.Time) bool {
	return reservation.Status == entities.Approved &&
		now.After(reservation.ReservationStart) &&
		now.Before(reservation.ReservationEnd)
}
