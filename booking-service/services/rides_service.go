package services

import (
	"context"
	"log"

	"github.com/essaubaid/ride-hailing/booking-service/handlers"
	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/proto/rides"
)

type RidesService struct {
	rides.UnimplementedRidesServiceServer
	handler handlers.RidesHandler
}

func NewRidesService(handler handlers.RidesHandler) *RidesService {
	return &RidesService{
		handler: handler,
	}
}

func (s *RidesService) UpdateRide(ctx context.Context, req *rides.UpdateRideRequest) (*rides.UpdateRideResponse, error) {
	log.Printf("gRPC: Received UpdateRide request for ID: %d", req.Id)

	ride := models.Ride{
		Id:          req.Id,
		Source:      req.Ride.Source,
		Destination: req.Ride.Destination,
		Distance:    req.Ride.Distance,
		Cost:        req.Ride.Cost,
	}

	_, err := s.handler.UpdateRide(ctx, &ride)
	if err != nil {
		return nil, err
	}

	return &rides.UpdateRideResponse{
		Message: "Ride updated successfully",
	}, nil
}

func (s *RidesService) CreateRide(ctx context.Context, req *rides.CreateRideRequest) (*rides.CreateRideResponse, error) {
	log.Printf("gRPC: Received CreateRide request for Source: %s, Destination: %s", req.Ride.Source, req.Ride.Destination)

	ride := models.Ride{
		Source:      req.Ride.Source,
		Destination: req.Ride.Destination,
		Distance:    req.Ride.Distance,
		Cost:        req.Ride.Cost,
	}

	_, err := s.handler.CreateRide(ctx, &ride)
	if err != nil {
		return nil, err
	}

	return &rides.CreateRideResponse{
		Message: "Ride created successfully",
	}, nil
}
