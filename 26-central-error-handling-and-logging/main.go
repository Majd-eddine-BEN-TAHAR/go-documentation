package main

import (
	"fmt"
	"net/http"

	"errors_handling/pkg/errors" // Importing custom errors package for API error handling.

	"go.uber.org/zap" // Importing zap for structured logging.
)

var logger *zap.Logger // Global logger instance.

func main() {
	var err error
	// Initializing the zap logger for structured logging.
	logger, err = zap.NewProduction()
	if err != nil {
		panic("Cannot initialize zap logger") // Panic if logger initialization fails.
	}
	defer logger.Sync() // Ensuring that logs are flushed at the end.

	// Setting up HTTP routes for the testing handlers.
	http.HandleFunc("/test", testDatabaseHandler)
	http.HandleFunc("/test2", testValidationHandler)
	
	// Starting the HTTP server on port 8080.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err)) // Logging fatal errors if server fails to start.
	}
}

func testDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	// Simulating a database error.
    simulatedDBError := fmt.Errorf("failed to connect to database")

    // Use NewDatabaseError to create an APIError
    apiErr := errors.NewDatabaseError(simulatedDBError, "Database connection error")
	// Handling the error: logging it and sending a structured JSON response.
    errors.HandleError(w, logger, apiErr)
}


func testValidationHandler(w http.ResponseWriter, r *http.Request) {
	// Example usage of the NewValidationError function.
	// Simulating a validation error.
	simulatedValidationError := fmt.Errorf("invalid input data")

	// Creating an API error for the validation error scenario.
	apiValidationError := errors.NewValidationError(simulatedValidationError, "Input validation failed")
	// Handling the error: logging it and sending a structured JSON response.
	errors.HandleError(w, logger, apiValidationError)
}