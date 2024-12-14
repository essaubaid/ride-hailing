package repositories

import (
	"database/sql"

	"github.com/essaubaid/ride-hailing/common/models"
)

type RidesRepository struct {
	db *sql.DB
}

func NewRidesRepository(db *sql.DB) *RidesRepository {
	return &RidesRepository{db: db}
}

func (r *RidesRepository) UpdateRideById(id int32, ride *models.Ride) error {
	_, err := r.db.Exec("UPDATE rides SET source = $1, destination = $2, distance = $3, cost = $4 WHERE id = $5",
		ride.Source,
		ride.Destination,
		ride.Distance,
		ride.Cost,
		id,
	)

	return err
}

func (r *RidesRepository) CreateRide(ride *models.Ride) error {
	err := r.db.QueryRow("INSERT INTO rides (source, destination, distance, cost) VALUES ($1, $2, $3, $4) RETURNING id",
		ride.Source,
		ride.Destination,
		ride.Distance,
		ride.Cost,
	).Scan(&ride.Id)

	return err
}
