package services

import (
	"errors"

	"database/sql"
	"event_booking_api/pkg/database"
)

// RegisterUserForEvent registers a user for a specific event.
func RegisterUserForEvent(userID, eventID int) error {
    // Check if the booking already exists to prevent duplicates.
    if BookingExists(userID, eventID) {
        return errors.New("user already registered for this event")
    }

    // Insert new booking
    query := `INSERT INTO bookings (user_id, event_id) VALUES (?, ?)`
    _, err := database.DB.Exec(query, userID, eventID)
    if err != nil {
		// Return an error if the insertion fails.
        return errors.New("failed to insert new booking into the database")
    }

    return nil
}

// BookingExists checks if a booking already exists for a given user and event.
func BookingExists(userID, eventID int) bool {
    var id int
    query := `SELECT id FROM bookings WHERE user_id = ? AND event_id = ?`
    err := database.DB.QueryRow(query, userID, eventID).Scan(&id)
    
	// Return true if a record is found, false if not.
    return err != sql.ErrNoRows
}


func CancelBooking(userID, eventID int) error {
	query := `DELETE FROM bookings WHERE user_id = ? and event_id = ?`
    _, err := database.DB.Exec(query, userID, eventID)
    return err
}



