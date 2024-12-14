package repositories

import (
	"database/sql"

	"github.com/essaubaid/ride-hailing/booking-service/responses"
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

func (r *BookingRepository) GetBooking(id int32) (*responses.BookingDetails, error) {
	bookingDetails := &responses.BookingDetails{}

	err := r.db.QueryRow(`SELECT u.name, r.source, r.destination, r.distance, r.cost, b.time 
		FROM bookings b 
		INNER JOIN users u ON b.user_id = u.id 
		INNER JOIN rides r ON b.ride_id = r.id 
		WHERE b.id = $1`, id).Scan(
		&bookingDetails.Name,
		&bookingDetails.Source,
		&bookingDetails.Destination,
		&bookingDetails.Distance,
		&bookingDetails.Cost,
		&bookingDetails.Time,
	)
	if err != nil {
		return nil, err
	}

	return bookingDetails, nil
}
