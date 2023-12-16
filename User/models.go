package user

import "time"

type User struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Mobile        string    `json:"mobile"`
	Email         string    `json:"email"`
	DriverLicense string    `json:"driver_license,omitempty"`
	CarPlate      string    `json:"car_plate,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

// Implement database-related functions here
