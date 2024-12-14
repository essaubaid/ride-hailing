package handlers

import (
	"context"
	"time"

	"github.com/essaubaid/ride-hailing/booking-service/repositories"
	"github.com/essaubaid/ride-hailing/booking-service/responses"
	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/proto/rides"
)

type BookingHandler struct {
	repo       *repositories.BookingRepository
	rideClient *rides.RidesServiceClient
}

func NewBookingHandler(repo *repositories.BookingRepository, rideClient *rides.RidesServiceClient) *BookingHandler {
	return &BookingHandler{
		repo:       repo,
		rideClient: rideClient,
	}
}

func (h *BookingHandler) GetBooking(ctx context.Context, bookingId int32) (*responses.BookingDetails, error) {

	bookingDetails, err := h.repo.GetBooking(bookingId)
	if err != nil {
		return nil, err
	}

	return bookingDetails, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, userId int32, ride models.Ride) (*models.Booking, error) {

	rideReq := &rides.CreateRideRequest{
		Ride: &rides.RideDetails{
			Source:      ride.Source,
			Destination: ride.Destination,
			Distance:    ride.Distance,
			Cost:        ride.Cost,
		},
	}

	rideResp, err := (*h.rideClient).CreateRide(ctx, rideReq)
	if err != nil {
		return nil, err
	}

	booking := &models.Booking{
		UserId: userId,
		RideId: rideResp.Id,
		Time:   time.Now(),
	}

	err = h.repo.CreateBooking(booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
