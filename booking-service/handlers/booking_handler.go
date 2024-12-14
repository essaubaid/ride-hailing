package handlers

import (
	"context"
	"time"

	"github.com/essaubaid/ride-hailing/common/models"
)

type Booking struct {
	Id     int32
	UserId int32
	RideId int32
	Time   time.Time
}

type BookingHandler struct {
}

func NewBookingHandler() *BookingHandler {
	return &BookingHandler{}
}

func (h *BookingHandler) GetBooking(ctx context.Context, bookingId int32) (*Booking, *models.Ride, error) {

	return &Booking{
			Id:     bookingId,
			UserId: 1,
			RideId: 1,
			Time:   time.Now(),
		}, &models.Ride{
			Id:          1,
			Source:      "Source",
			Destination: "Destination",
			Distance:    10,
			Cost:        100,
		}, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, userId int32, ride models.Ride) (*Booking, error) {

	return &Booking{
		Id:     1,
		UserId: userId,
		RideId: 1,
		Time:   time.Now(),
	}, nil
}
