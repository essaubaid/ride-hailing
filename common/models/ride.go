package models

type Ride struct {
	Id          int32  `json:"id"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Distance    int32  `json:"distance"`
	Cost        int32  `json:"cost"`
}
