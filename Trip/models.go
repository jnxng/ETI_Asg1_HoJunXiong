package trip

import "time"

type Trip struct {
	ID                int       `json:"id"`
	CarOwnerID        int       `json:"car_owner_id"`
	PickupLocation    string    `json:"pickup_location"`
	AltPickupLocation string    `json:"alt_pickup_location,omitempty"`
	StartTime         time.Time `json:"start_time"`
	Destination       string    `json:"destination"`
	AvailableSeats    int       `json:"available_seats"`
}

// Implement database-related functions here
