package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "jnxng:Asg1@tcp(127.0.0.1:3306)/carpooling")
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(user *User) error {
	query := "INSERT INTO users (first_name, last_name, mobile, email, driver_license, car_plate, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, user.FirstName, user.LastName, user.Mobile, user.Email, user.DriverLicense, user.CarPlate, user.CreatedAt)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err
	}

	user.ID = int(lastInsertID)
	return nil
}

func updateUser(user *User) error {
	query := "UPDATE users SET first_name=?, last_name=?, mobile=?, email=?, driver_license=?, car_plate=? WHERE id=?"

	_, err := db.Exec(query, user.FirstName, user.LastName, user.Mobile, user.Email, user.DriverLicense, user.CarPlate, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func deleteUser(userID int) error {
	query := "DELETE FROM users WHERE id=?"

	_, err := db.Exec(query, userID)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	return nil
}

func getUserTrips(userID int) ([]Trip, error) {
	query := "SELECT * FROM trips WHERE car_owner_id=?"

	trips, err := db.Exec(query, userID)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	
	return trips, nil
}