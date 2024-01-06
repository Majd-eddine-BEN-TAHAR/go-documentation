package handlers

import (
	"encoding/json"
	"event_booking_api/internal/app/middlewares"
	"event_booking_api/internal/app/services"
	"event_booking_api/pkg/errors"
	"net/http"
)

// - POST for creating a booking.
// - DELETE for canceling a booking.
func BookingHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method {
		case http.MethodPost:
			middlewares.AuthMiddleware(handleEventBooking)(w, r)
		case http.MethodDelete:
			middlewares.AuthMiddleware(handleCancelBooking)(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleEventBooking(w http.ResponseWriter,r *http.Request){
    // Extract event ID from the URL
	eventID, _ := extractEventID(r) // Error handling have been be done in the router
    
	// Extract user ID from the context (set by auth middleware)
    userID := r.Context().Value("userID").(int)

	// Validate that the event exists
    event, err := services.GetEventByID(eventID)
    if err != nil || event == nil {
        errors.HandleError(w, errors.NewAPIError(nil, "Event not found", http.StatusNotFound))
        return
    }

	// Perform the booking
	if err := services.RegisterUserForEvent(userID, eventID); err != nil {
		errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusInternalServerError))
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
	 Message string `json:"message"`
	}{"Registered for event successfully"})
}
	
func handleCancelBooking(w http.ResponseWriter, r *http.Request) {
    // Extract event ID from the URL
	eventID, _ := extractEventID(r) // Error handling have been be done in the router
    
	// Extract user ID from the context (set by auth middleware)
    userID := r.Context().Value("userID").(int)

	// Check if the event exists
    event, err := services.GetEventByID(eventID)
    if err != nil || event == nil {
        errors.HandleError(w, errors.NewAPIError(nil, "Event not found", http.StatusNotFound))
        return
    }

	// Check if the user is registered for the event
    if !services.BookingExists(userID, eventID) {
		errors.HandleError(w, errors.NewAPIError(nil, "User is not registered for this event", http.StatusNotFound))
		return
    }

    // Unregister the user from the event
    if err := services.CancelBooking(userID,eventID); err != nil {
		errors.HandleError(w, errors.NewAPIError(err, "Error when deleting the booking", http.StatusInternalServerError))
		return
	};
	
    // Respond with HTTP status code 204 - No Content
	w.WriteHeader(http.StatusNoContent)
}
