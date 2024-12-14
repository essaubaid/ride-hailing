package responses

import "time"

type BookingDetails struct {
	Name        string    `json:"name"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Distance    int32     `json:"distance"`
	Cost        int32     `json:"cost"`
	Time        time.Time `json:"time"`
}
