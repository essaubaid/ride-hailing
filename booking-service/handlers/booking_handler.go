package handlers

import (
	"context"
	"time"

	"github.com/essaubaid/ride-hailing/booking-service/repositories"
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

func (h *BookingHandler) GetBooking(ctx context.Context, bookingId int32) (*models.Booking, *models.Ride, error) {

	return &models.Booking{
			Id:     1,
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
