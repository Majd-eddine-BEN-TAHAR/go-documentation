package services

import (
	"errors"
	"event_booking_api/internal/app/models"
	"event_booking_api/internal/app/validations"
	"event_booking_api/pkg/auth"
	"event_booking_api/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles the registration of a new user.
// It hashes the user's password and stores the user's details in the database.
func RegisterUser(user models.User) error {
	// Validate user input
    if err := validations.ValidateUser(user); err != nil {
        return err
    }

    // Check if the email already exists in the database
    var emailCount int
    err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&emailCount)
    if err != nil {
        return err
    }

    if emailCount > 0 {
        return errors.New("email already exists")
    }

    // Check if the username already exists in the database
    var usernameCount int
    err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&usernameCount)
    if err != nil {
        return err
    }

    if usernameCount > 0 {
        return errors.New("username already exists")
    }

    // Hash the user's password using bcrypt.
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err // Return error if password hashing fails
    }

    // Insert the new user into the database.
    // The hashed password is stored, not the plain one.
    _, err = database.DB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
        user.Username, string(hashedPassword), user.Email)
    return err // Return error if the database operation fails
}


// AuthenticateUser checks the credentials and returns a JWT token if they are valid
func AuthenticateUser(creds models.Credentials) (string, error) {
    
	// Validate user input
	if err := validations.ValidateLogin(creds); err != nil {
		return "", err
	}

	// Retrieve user from database by username
    var user models.User
    err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", creds.Username).Scan(&user.ID,&user.Username, &user.Password)
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    // Compare the provided password with the stored hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        return "", errors.New("invalid username or password")
    }

    // Generate and return a JWT token
    token, err := auth.GenerateToken(user.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}