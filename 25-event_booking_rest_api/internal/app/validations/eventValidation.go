package validations

import (
	"errors"
	"event_booking_api/internal/app/models"
	"time"
)

// ValidateEvent checks the validity of the event data.
func ValidateEvent(event models.Event) error {
    if len(event.Title) < 4 || len(event.Description) < 10 {
        return errors.New("title must be at least 4 characters and description at least 10 characters long")
    }

    // Check if the location is a non-empty string
    if len(event.Location) == 0 {
        return errors.New("event location must be specified")
    }

    // Validate that the start time is before the end time
    if !event.StartTime.Before(event.EndTime) {
        return errors.New("start time must be before end time")
    }

    // Ensure the event start time is not in the past
    if event.StartTime.Before(time.Now()) {
        return errors.New("event start time must be in the future")
    }

    return nil
}