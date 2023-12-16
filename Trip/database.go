package trip

import (
	"database/sql"
	"log"
	"time"

	"github.com/jnxng/ETI_Asg1_HoJunXiong/common"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "jnxng:asg1@tcp(127.0.0.1:3306)/carpooling")
	if err != nil {
		log.Fatal(err)
	}
}

func createTrip(trip *Trip) error {
	query := "INSERT INTO trips (car_owner_id, pickup_location, alt_pickup_location, start_time, destination, available_seats) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, trip.CarOwnerID, trip.PickupLocation, trip.AltPickupLocation, trip.StartTime, trip.Destination, trip.AvailableSeats)
	if err != nil {
		log.Println("Error creating trip:", err)
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err
	}

	trip.ID = int(lastInsertID)
	return nil
}

func startTrip(tripID int) error {
	query := "UPDATE trips SET started_at=? WHERE id=?"

	_, err := db.Exec(query, time.Now(), tripID)
	if err != nil {
		log.Println("Error starting trip:", err)
		return err
	}

	return nil
}

func cancelTrip(tripID int) error {
	query := "DELETE FROM trips WHERE id=?"

	_, err := db.Exec(query, tripID)
	if err != nil {
		log.Println("Error canceling trip:", err)
		return err
	}

	return nil
}