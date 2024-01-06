package validations

import (
	"errors"

	"event_booking_api/internal/app/models"
)

// ValidateUser validates the user's data
func ValidateUser(user models.User) error {
    if len(user.Username) < 4 || len(user.Password) < 4 {
        return errors.New("username and password must be at least 4 characters long")
    }

    if !isValidEmail(user.Email) {
        return errors.New("invalid email address")
    }

    return nil
}


// ValidateLogin validates the login credentials
func ValidateLogin(creds models.Credentials) error {
    // Check if username and password are provided and meet basic length requirements
    if len(creds.Username) < 4 || len(creds.Password) < 4 {
        return errors.New("username and password must be at least 4 characters long")
    }

    // Add any other specific login validations here if needed

    return nil
}
