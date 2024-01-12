package util

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// ValidationErrorResponse represents the structured error response
type ValidationErrorResponse struct {
    Error       string       `json:"error"`
    Message     string       `json:"message"`
    FieldErrors []FieldError `json:"field_errors"`
}

// FieldError represents the detailed error for a specific field
type FieldError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// HandleValidationError processes validation errors and sends a structured response
// It expects the validated object as an argument to fetch custom error messages from struct tags
func HandleValidationError(w http.ResponseWriter, err error, obj interface{}) {
    if ve, isValidationError := err.(validator.ValidationErrors); isValidationError {
        var fieldErrors []FieldError

        // Use reflection to get the type of the struct
        objType := reflect.TypeOf(obj)
        if objType.Kind() == reflect.Ptr {
            objType = objType.Elem()
        }

        for _, fe := range ve {
            field := fe.StructField()

            // Find the field in the struct
            fieldType, exists := objType.FieldByName(field)
            if !exists {
                continue
            }

            // Get the custom error message from the tag, default to generic message if not found
            customMsg, ok := fieldType.Tag.Lookup("errMsg")
            if !ok {
                customMsg = "Invalid value for " + field + ", failed on " + fe.Tag()
            }

            fieldError := FieldError{
                Field:   field,
                Message: customMsg,
            }
            fieldErrors = append(fieldErrors, fieldError)
        }

        errorResponse := ValidationErrorResponse{
            Error:       "Validation failed",
            Message:     "Input validation errors",
            FieldErrors: fieldErrors,
        }

        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(errorResponse)
    } else {
        // Handle non-validation errors
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}