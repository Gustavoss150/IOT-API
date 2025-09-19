package reservationRepo

import "api/entities"

type ReservationRepository interface {
	Save(reservation *entities.Reservation) error
}
