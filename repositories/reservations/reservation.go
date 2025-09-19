package reservationRepo

import (
	"api/config"
	"api/entities"
	"errors"

	"gorm.io/gorm"
)

type reservationRepository struct {
	DB *gorm.DB
}

func InitReservationDatabase() (ReservationRepository, error) {
	db := config.DB
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}
	return &reservationRepository{DB: db}, nil
}

func (r *reservationRepository) Save(reservation *entities.Reservation) error {
	return r.DB.Save(reservation).Error
}
