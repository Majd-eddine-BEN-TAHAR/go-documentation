package handlers

import (
	"encoding/json"
	"net/http"

	"event_booking_api/internal/app/middlewares"
	"event_booking_api/internal/app/models"
	"event_booking_api/internal/app/services"
	"event_booking_api/pkg/errors"
)

// - GET to fetch all events.
// - POST to create a new event (with authentication).
func EventsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            // Fetch all events
            handleGetEvent(w, r)
        case http.MethodPost:
            middlewares.AuthMiddleware(handlePostEvent)(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGetEvent(w http.ResponseWriter, r *http.Request){
    // Fetch upcoming events
    
    events, err := services.GetUpcomingEvents()
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusInternalServerError))
        return
    }

    // Respond with the list of events
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(events)
}

// handlePostEvent handles the event creation request.
func handlePostEvent(w http.ResponseWriter, r *http.Request) {
    var event models.Event
	
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid request body", http.StatusBadRequest))
        return
    }

    // Extract the user ID from the request context
    userID := r.Context().Value("userID").(int)

    if err := services.CreateEvent(event, userID); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusBadRequest))
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{"Event created successfully"})
}