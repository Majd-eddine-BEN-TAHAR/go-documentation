package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError represents an error with an additional status code
type APIError struct {
    Error   error
    Message string
    Code    int
}

// NewAPIError creates a new APIError
func NewAPIError(err error, message string, code int) *APIError {
    return &APIError{
        Error:   err,
        Message: message,
        Code:    code,
    }
}

// HandleError handles the error by logging it and sending an appropriate response to the client
func HandleError(w http.ResponseWriter, err *APIError) {
    log.Printf("Error: %v", err.Error) // Log the error

    // Send a JSON response with the error message and status code
    w.WriteHeader(err.Code)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{
        Message: err.Message,
    })
}