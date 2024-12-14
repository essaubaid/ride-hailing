package models

import "time"

type Booking struct {
	Id     int32     `json:"id"`
	UserId int32     `json:"user_id"`
	RideId int32     `json:"ride_id"`
	Time   time.Time `json:"time"`
}
