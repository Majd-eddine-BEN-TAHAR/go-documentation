package main

import (
	"errors"
	"fmt"
)

// Example function that returns an error
func mightFail() (string, error) {
	// Example condition where an error might occur
	failureCondition := true

	if failureCondition {
		// Returning an error using the errors package provided by go
		return "", errors.New("something went wrong")
	}
	// Normal return when no error occurs
	return "success", nil
}

func main() {
	// Calling a function that might fail
	result, err := mightFail()

	// Explicit Error Handling: Check if there is an error
	if err != nil {
		// Handling the error, e.g., logging or deciding what to do next
		fmt.Println("Error occurred:", err)
		// Optionally, handle the error based on its type or content
		// if err == someSpecificError {
		//     // specific error handling
		// }

		// Errors don't crash your app; you control the flow
		// Decide whether to continue, retry, or stop
	} else {
		// If no error, normal execution continues
		fmt.Println("Operation successful:", result)
	}

	// Calling a function that will return a custom error
	customErr := customError()

	// Check if the function returned a custom error
	if customErr != nil {
		// If an error is returned, print its message.
		// Since customErr is of type CustomErrorType, it has the 'Error' method that returns the message.
		fmt.Println("Custom Error occurred:", customErr.Error())
	
		// You can also perform specific actions based on custom error properties.
		// For example, you might want to check the error message and decide what to do.
		if customErr.Error() == "a specific error message" {
			// Handle this specific error case
		}
	}
}

// Custom error creation and handling
// Define a custom error type
// CustomErrorType is our own error type that implements the error interface.
type CustomErrorType struct {
	Message string // Message will store the error message
}

// Error implements the error interface for CustomErrorType
// It returns the error message as a string.
// Error is a method that makes CustomErrorType satisfy the error interface.
func (e *CustomErrorType) Error() string {
	return e.Message
}

// Function that demonstrates returning our custom error
func customError() error {
	// Here, we're simulating a condition where we need to return a custom error.
	// We create an instance of CustomErrorType with a specific message.
	return &CustomErrorType{"a custom error occurred"}
}