package reservationRepo

import (
	"api/entities"
	"time"
)

type ReservationRepository interface {
	Save(reservation *entities.Reservation) error
	GetByID(id string) (*entities.Reservation, error)
	HasReservationConflict(machineID string, start, end time.Time) (bool, error)
	GetAllReservationsByDay(day time.Time) ([]entities.Reservation, error)
}
