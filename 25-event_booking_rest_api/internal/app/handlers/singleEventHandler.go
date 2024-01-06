package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
    // Set a limit for the size of the form data (10 MB).
    const maxUploadSize = 10 << 20 // 10 MB
    // Parse the multipart form with the specified maximum size.
    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        // Handle errors related to form size limits.
        errors.HandleError(w, errors.NewAPIError(err, "File too large", http.StatusBadRequest))
        return
    }

    // Extract the eventID from the URL path.
    eventID, err := extractEventID(r)
    if err != nil {
        // Handle errors if the eventID is not valid or not found.
        errors.HandleError(w, errors.NewAPIError(err, "Invalid event ID", http.StatusBadRequest))
        return
    }

    // Attempt to decode the JSON event data from the form.
    var updatedEvent models.Event
    err = json.Unmarshal([]byte(r.FormValue("event")), &updatedEvent)
    if err != nil {
        // Handle errors if the event data is not valid JSON.
        errors.HandleError(w, errors.NewAPIError(err, "Invalid event data", http.StatusBadRequest))
        return
    }

    // Retrieve the user ID from the request context, assuming it was set earlier (e.g., by an authentication middleware).
    userID := r.Context().Value("userID").(int)

    // Retrieve the existing event to check if it exists and if the current user is the creator.
    existingEvent, err := services.GetEventByID(eventID)
    if err != nil {
        // Handle errors if the event is not found.
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusNotFound))
        return
    }

    // Check if the user is authorized to update the event.
    if existingEvent.CreatorID != userID {
        // Return an unauthorized error if the user is not the event creator.
        errors.HandleError(w, errors.NewAPIError(nil, "Unauthorized to update this event", http.StatusUnauthorized))
        return
    }

    // Handle the file upload part, if an image file is provided.
    file, header, err := r.FormFile("image")
    if err == nil {
        defer file.Close()

        // Check if the uploaded file size is within the allowed limit.
        if header.Size > maxUploadSize {
            // Handle errors related to file size limit.
            errors.HandleError(w, errors.NewAPIError(nil, "File size exceeds the limit", http.StatusBadRequest))
            return
        }

        // Validate the MIME type of the file (only allow JPEG and PNG).
        allowedMIMEs := map[string]bool{"image/jpeg": true, "image/jpg": true, "image/png": true}
        if _, allowed := allowedMIMEs[header.Header.Get("Content-Type")]; !allowed {
            // Handle errors if the file type is not allowed.
            errors.HandleError(w, errors.NewAPIError(nil, "Invalid file type", http.StatusBadRequest))
            return
        }

        // Create a new file in the server's filesystem to save the uploaded file.
        fileName := "event_" + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(header.Filename)
        dstPath := filepath.Join("uploads", fileName)
        dst, err := os.Create(dstPath)
        if err != nil {
            // Handle errors related to file creation on the server.
            errors.HandleError(w, errors.NewAPIError(err, "Error creating file", http.StatusInternalServerError))
            return
        }
        defer dst.Close()

        // Copy the uploaded file data to the server's file.
        if _, err := io.Copy(dst, file); err != nil {
            // Handle errors related to writing the file data.
            errors.HandleError(w, errors.NewAPIError(err, "Error saving file", http.StatusInternalServerError))
            return
        }

        // Update the image URL in the event with the path of the new uploaded image.
        updatedEvent.ImageURL = dstPath
    } else {
        // If no new image is provided, keep the existing image URL.
        updatedEvent.ImageURL = existingEvent.ImageURL
    }

    // Validate the updated event data (excluding image URL).
    if err := validations.ValidateEvent(updatedEvent); err != nil {
        // Handle errors if the updated event data is not valid.
        errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusBadRequest))
        return
    }

    // Update the event in the database with the new data.
    if err := services.UpdateEvent(eventID, updatedEvent); err != nil {
        // Handle errors related to the database update operation.
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