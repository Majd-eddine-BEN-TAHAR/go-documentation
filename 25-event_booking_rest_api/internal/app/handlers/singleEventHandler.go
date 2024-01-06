package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"event_booking_api/internal/app/middlewares"
	"event_booking_api/internal/app/models"
	"event_booking_api/internal/app/services"
	"event_booking_api/internal/app/validations"
	"event_booking_api/pkg/errors"
)

// - GET to fetch a single event.
// - PUT to update an event (with authentication).
// - DELETE to delete an event (with authentication).
func SingleEventHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            // Fetch single event
            handleGetSingleEvent(w, r)
		case http.MethodPut:
            // update single event
			middlewares.AuthMiddleware(handlePutSingleEvent)(w, r)
				
        case http.MethodDelete:
            // delete sinle event
            middlewares.AuthMiddleware(handleDeleteSingleEvent)(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


func handleGetSingleEvent(w http.ResponseWriter, r *http.Request){
	// extract eventID from the URL
	eventID, err := extractEventID(r)
	if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid event ID", http.StatusBadRequest))
        return
    }

	// Fetch upcoming events
    event, err := services.GetEventByID(eventID)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusInternalServerError))
        return
    }

    // Respond with the list of events
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(event)
}

func handlePutSingleEvent(w http.ResponseWriter, r *http.Request) {
    // Decode the incoming JSON request to the Event struct.
    var updatedEvent models.Event
    if err := json.NewDecoder(r.Body).Decode(&updatedEvent); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid request body", http.StatusBadRequest))
        return
    }

	// extract eventID from the URL
	eventID, err := extractEventID(r)
	if err != nil {
		errors.HandleError(w, errors.NewAPIError(err, "Invalid event ID", http.StatusBadRequest))
		return
	}
	
    // Validate the updated event
    if err := validations.ValidateEvent(updatedEvent); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusBadRequest))
        return
    }
	
    // get the userID from the context
	userID := r.Context().Value("userID").(int)

    // Check if the event exists and if the user is the creator
    existingEvent, err := services.GetEventByID(eventID)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusNotFound))
        return
    }

    if existingEvent.CreatorID != userID {
        errors.HandleError(w, errors.NewAPIError(nil, "Unauthorized to update this event", http.StatusUnauthorized))
        return
    }

    // Update the event
    if err := services.UpdateEvent(eventID, updatedEvent); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusInternalServerError))
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{"Event updated successfully"})
}

func handleDeleteSingleEvent(w http.ResponseWriter, r *http.Request) {
    // Extract eventID from the URL
    eventID, err := extractEventID(r)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid event ID", http.StatusBadRequest))
        return
    }

    // Get the userID from the context
    userID := r.Context().Value("userID").(int)

    // Check if the event exists and if the user is the creator
    event, err := services.GetEventByID(eventID)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Event not found", http.StatusNotFound))
        return
    }

    if event.CreatorID != userID {
        errors.HandleError(w, errors.NewAPIError(nil, "Unauthorized to delete this event", http.StatusUnauthorized))
        return
    }

    // Delete the event
    if err := services.DeleteEvent(eventID); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Error deleting event", http.StatusInternalServerError))
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{"Event deleted successfully"})
}

// Function to extract eventID from the URL
func extractEventID(r *http.Request) (int, error) {
    path := strings.Trim(r.URL.Path, "/")
    pathSegments := strings.Split(path, "/")
    // Assuming the eventID is the second segment in the URL
    eventIdStr := pathSegments[1]
    eventID, err := strconv.Atoi(eventIdStr)
    if err != nil {
        return 0, err // Return an error if conversion fails
    }
    return eventID, nil
}