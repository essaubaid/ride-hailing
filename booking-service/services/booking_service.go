package services

import (
	"context"
	"log"

	"github.com/essaubaid/ride-hailing/booking-service/handlers"
	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/proto/booking"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingService struct {
	booking.UnimplementedBookingServiceServer
	handler handlers.BookingHandler
}

func NewBookingService(handler handlers.BookingHandler) *BookingService {
	return &BookingService{
		handler: handler,
	}
}

func (s *BookingService) GetBooking(ctx context.Context, req *booking.GetBookingRequest) (*booking.GetBookingResponse, error) {
	log.Printf("gRPC: Received GetBooking request for Booking ID: %d", req.Id)

	bookingData, rideData, err := s.handler.GetBooking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &booking.GetBookingResponse{
		Name:        "User Name",
		Source:      rideData.Source,
		Destination: rideData.Destination,
		Distance:    rideData.Distance,
		Cost:        rideData.Cost,
		Time:        timestamppb.New(bookingData.Time),
	}, nil
}

func (s *BookingService) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {
	log.Printf("gRPC: Received CreateBooking request for User ID: %d", req.UserId)

	ride := models.Ride{
		Source:      req.Ride.Source,
		Destination: req.Ride.Destination,
		Distance:    req.Ride.Distance,
		Cost:        req.Ride.Cost,
	}
	bookingData, err := s.handler.CreateBooking(ctx, req.UserId, ride)
	if err != nil {
		return nil, err
	}

	return &booking.CreateBookingResponse{
		Booking: &booking.BookingDetails{
			UserId: bookingData.UserId,
			RideId: bookingData.RideId,
			Time:   timestamppb.New(bookingData.Time),
		},
	}, nil
}
