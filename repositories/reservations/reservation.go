package reservationRepo

import (
	"api/config"
	"api/entities"
	"errors"
	"time"

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

func (r *reservationRepository) GetByID(id string) (*entities.Reservation, error) {
	var reservation entities.Reservation
	if err := r.DB.Where("id = ?", id).First(&reservation).Error; err != nil {
		return nil, errors.New("reserva não encontrada")
	}
	return &reservation, nil
}

func (r *reservationRepository) HasReservationConflict(equipmentID string, start, end time.Time) (bool, error) {
	var count int64
	err := r.DB.Model(&entities.Reservation{}).
		Where("equipment_id = ? AND reservation_end > ? AND reservation_start < ?", equipmentID, start, end).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *reservationRepository) GetAllReservationsByDay(day time.Time) ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	if err := r.DB.Where("reservation_start >= ? AND reservation_start < ?", startOfDay, endOfDay).Find(&reservations).Error; err != nil {
		return nil, errors.New("sem reservas encontradas para o dia especificado")
	}
	return reservations, nil
}
