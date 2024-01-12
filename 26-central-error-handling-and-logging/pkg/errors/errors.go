package errors

import (
	"encoding/json" // Used for encoding error messages into JSON.
	"net/http"

	"go.uber.org/zap" // A logging library for structured logging.
)

// Defining constants for common error types.
const (
	ErrTypeDatabase   = "DatabaseError"   // Represents errors related to the database.
	ErrTypeValidation = "ValidationError" // Represents validation errors.
)

// APIError struct defines the structure of an API error.
type APIError struct {
	Err     error  // The original error.
	Message string // Human-readable message describing the error.
	Code    int    // HTTP status code associated with this error.
	Type    string // Type of error (e.g., database, validation).
}

// NewAPIError is a constructor for creating a new APIError.
// Parameters: Original error, custom message, HTTP status code, and error type.
// Returns a pointer to a new APIError instance.
func NewAPIError(err error, message string, code int, errType string) *APIError {
	return &APIError{
		Err:     err,
		Message: message,
		Code:    code,
		Type:    errType,
	}
}

// HandleError is responsible for handling an API error.
// It logs the error and sends a JSON response with the error details.
func HandleError(w http.ResponseWriter, logger *zap.Logger, apiErr *APIError) {
	// Logging the error using zap logger.
	logger.Error("API error occurred",
		zap.Error(apiErr.Err),
		zap.String("type", apiErr.Type),
	)

	// Setting up the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")
	// Writing the HTTP status code.
	w.WriteHeader(apiErr.Code)
	// Encoding and sending the error message in JSON format.
	json.NewEncoder(w).Encode(map[string]string{"message": apiErr.Message, "type": apiErr.Type})
}

/* Utility functions for creating specific API errors */

// NewDatabaseError creates a new APIError specifically for database errors.
func NewDatabaseError(err error, message string) *APIError {
	return NewAPIError(err, message, http.StatusInternalServerError, ErrTypeDatabase)
}

// NewValidationError creates a new APIError specifically for validation errors.
func NewValidationError(err error, message string) *APIError {
	return NewAPIError(err, message, http.StatusBadRequest, ErrTypeValidation)
}