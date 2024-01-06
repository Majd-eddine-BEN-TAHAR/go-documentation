package handlers

import (
	"encoding/json"
	"net/http"

	"event_booking_api/internal/app/models"
	"event_booking_api/internal/app/services"
	"event_booking_api/pkg/errors"
)

// RegisterHandler handles the user registration requests.
// It supports POST method to register a new user.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodPost:
            handlePostRegister(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// LoginHandler handles the user login requests.
// It supports POST method to authenticate a user.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodPost:
            handlePostLogin(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


// handlePostRegister is the HTTP post handler for the user registration endpoint.
// It processes the user registration request and responses.
func handlePostRegister(w http.ResponseWriter, r *http.Request) {
    var user models.User

    // Decode the incoming JSON request to the User struct and put in the user variable.
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid request body", http.StatusBadRequest))
        return
    }

    // Call the RegisterUser service to register the user.
    if err := services.RegisterUser(user); err != nil {
        // Pass the specific error message from the service
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusBadRequest))
        return // Stop processing and return an error if registration fails
    }

    // Respond with a success message if registration is successful.
	// Set the HTTP response status code to 201 Created.
	w.WriteHeader(http.StatusCreated)

	// Create a JSON encoder that writes to the http.ResponseWriter 'w'.
	encoder := json.NewEncoder(w)

	// Create an anonymous struct with a 'Message' field and initialize it
	// with the message indicating successful user registration.
	response := struct {
		Message string `json:"message"`
	}{"User registered successfully"}

	// Encode the 'response' struct into JSON format and send it as the
	// HTTP response body.
	encoder.Encode(response)
}


// handlePostLogin handles user login post requests
func handlePostLogin(w http.ResponseWriter, r *http.Request) {
    var creds models.Credentials 

    // Decode the incoming JSON request to the Credentials struct and put in the creds variable.
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid request body", http.StatusBadRequest))
        return
    }

    // Authenticate the user
	// err.Error() will add the error message returned from AuthenticateUser
    token, err := services.AuthenticateUser(creds)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusUnauthorized))
        return
    }

    // Respond with JWT token if login is successful
    json.NewEncoder(w).Encode(struct {
        Token string `json:"token"`
    }{Token: token})
}