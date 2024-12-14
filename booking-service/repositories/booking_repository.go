package repositories

import (
	"database/sql"

	"github.com/essaubaid/ride-hailing/common/models"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db}
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
	err := r.db.QueryRow("INSERT INTO bookings (user_id, ride_id, time) VALUES ($1, $2, $3) RETURNING id",
		booking.UserId,
		booking.RideId,
		booking.Time,
	).Scan(&booking.Id)
	return err
}
