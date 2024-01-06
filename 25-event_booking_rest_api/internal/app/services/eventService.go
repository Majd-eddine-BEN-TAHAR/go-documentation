package services

import (
	"event_booking_api/internal/app/models"
	"event_booking_api/internal/app/validations"

	"event_booking_api/pkg/database"
)

func CreateEvent(event models.Event) error {
    // Validate the event data before inserting into the database
    if err := validations.ValidateEvent(event); err != nil {
        return err
    }

    // Prepare SQL query to insert event data, including image URL
    query := `INSERT INTO events (title, description, location, start_time, end_time, creator_id, image_url)
    VALUES (?, ?, ?, ?, ?, ?, ?)`
    _, err := database.DB.Exec(query, event.Title, event.Description, event.Location, event.StartTime, event.EndTime, event.CreatorID, event.ImageURL)
    return err
}

// GetUpcomingEvents retrieves all events with a start time greater than the current time.
func GetUpcomingEvents() ([]models.Event, error) {
    var events []models.Event

    // Query to select events with start_time greater than the current time
    query := `SELECT id, title, description, location, start_time, end_time, creator_id, image_url FROM events WHERE start_time > CURRENT_TIMESTAMP`
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Iterating over the query results and appending them to the events slice.
    for rows.Next() {
        var event models.Event
        if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatorID, &event.ImageURL); err != nil {
            return nil, err
        }
        events = append(events, event)
    }

        // Checking for any errors occurred during iteration on rows.
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return events, nil
}

// GetEventByID retrieves an event by its ID from the database.
func GetEventByID(eventID int) (*models.Event, error) {
    var event models.Event

    // SQL query to retrieve the event data
    query := `SELECT * FROM events WHERE id = ?`
    
    // Executing the query with the event ID
    row := database.DB.QueryRow(query, eventID)

    // Scanning the query result into the event struct.
    err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatorID, &event.ImageURL)
    if err != nil {
        // Return the error if any occurred during query execution or putting values in the row
        return nil, err
    }

    // Return the retrieved event and nil error
    return &event, nil
}

// UpdateEvent updates an existing event's details in the database.
func UpdateEvent(id int, event models.Event) error {
    query := `SELECT id, title, description, location, start_time, end_time, image_url FROM events WHERE id = ?`
    _, err := database.DB.Exec(query, event.Title, event.Description, event.Location, event.StartTime, event.EndTime, event.ImageURL, id)
    return err
}

// DeleteEvent deletes an event from the database by its ID.
func DeleteEvent(eventID int) error {
    query := `DELETE FROM events WHERE id = ?`
    _, err := database.DB.Exec(query, eventID)
    return err
}
