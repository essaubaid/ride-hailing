package handlers

import (
	"context"

	"github.com/essaubaid/ride-hailing/booking-service/repositories"
	"github.com/essaubaid/ride-hailing/common/models"
)

type RidesHandler struct {
	repo *repositories.RidesRepository
}

func NewRidesHandler(repo *repositories.RidesRepository) *RidesHandler {
	return &RidesHandler{
		repo: repo,
	}
}

func (h *RidesHandler) UpdateRide(ctx context.Context, ride *models.Ride) (*models.Ride, error) {

	err := h.repo.UpdateRideById(ride.Id, ride)
	if err != nil {
		return nil, err
	}

	return ride, nil
}

func (h *RidesHandler) CreateRide(ctx context.Context, ride *models.Ride) (*models.Ride, error) {

	err := h.repo.CreateRide(ride)
	if err != nil {
		return nil, err
	}

	return ride, nil
}
