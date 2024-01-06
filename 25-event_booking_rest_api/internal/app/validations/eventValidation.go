package validations

import (
	"errors"
	"event_booking_api/internal/app/models"
	"os"
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

    // Validate the Image URL
    if err := validateImageURL(event.ImageURL); err != nil {
        return err
    }

    return nil
}

// validateImageURL checks if the provided image URL is valid.
func validateImageURL(imageURL string) error {
    // Check if the image URL is not empty
    if imageURL == "" {
        return errors.New("image URL must be provided")
    }

    // Check if the file exists at the given path
    if _, err := os.Stat(imageURL); os.IsNotExist(err) {
        return errors.New("image file does not exist at the provided path")
    }

    return nil
}