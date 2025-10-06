package reservationRepo

import (
	"api/entities"
	"time"
)

type ReservationRepository interface {
	Save(reservation *entities.Reservation) error
	GetByID(id string) (*entities.Reservation, error)
	HasReservationConflict(equipmentID string, start, end time.Time) (bool, error)
	GetAllReservationsByDay(day time.Time) ([]entities.Reservation, error)
	GetByUserID(userID string) ([]entities.Reservation, error)
	GetByMachineID(equipmentID string) ([]entities.Reservation, error)
}
