package handlers

import "context"

type Ride struct {
	Id          int32
	Source      string
	Destination string
	Distance    int32
	Cost        int32
}

type RidesHandler struct {
}

func NewRidesHandler() *RidesHandler {
	return &RidesHandler{}
}

func (h *RidesHandler) UpdateRide(ctx context.Context, ride Ride) (*Ride, error) {

	return &Ride{
		Id:          ride.Id,
		Source:      ride.Source,
		Destination: ride.Destination,
		Distance:    ride.Distance,
		Cost:        ride.Cost,
	}, nil
}
