package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
// func handlePostEvent(w http.ResponseWriter, r *http.Request) {
//     var event models.Event
	
//     if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
//         errors.HandleError(w, errors.NewAPIError(err, "Invalid request body", http.StatusBadRequest))
//         return
//     }

//     // Extract the user ID from the request context
//     userID := r.Context().Value("userID").(int)

//     if err := services.CreateEvent(event, userID); err != nil {
//         errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusBadRequest))
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(struct {
//         Message string `json:"message"`
//     }{"Event created successfully"})
// }

// handlePostEvent handles the creation of a new event, including image upload.
func handlePostEvent(w http.ResponseWriter, r *http.Request) {
    // Parse the multipart form with a maximum size (e.g., 10MB).
    const maxUploadSize = 10 << 20 // 10 MB
    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "File too large", http.StatusBadRequest))
        return
    }

    // Decode event data from the form value.
    var event models.Event
    err := json.Unmarshal([]byte(r.FormValue("event")), &event)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Invalid event data", http.StatusBadRequest))
        return
    }

    // Handle file upload.
    file, header, err := r.FormFile("image")
    if err != nil {
        if err == http.ErrMissingFile {
            errors.HandleError(w, errors.NewAPIError(nil, "No file uploaded", http.StatusBadRequest))
        } else {
            errors.HandleError(w, errors.NewAPIError(err, "Error retrieving the file", http.StatusBadRequest))
        }
        return
    }
    defer file.Close()

    // Validate file size.
    if header.Size > maxUploadSize {
        errors.HandleError(w, errors.NewAPIError(nil, "File size exceeds the limit", http.StatusBadRequest))
        return
    }

    // Validate file type (e.g., only allow JPEG and PNG).
    allowedMIMEs := map[string]bool{"image/jpeg": true, "image/jpg": true, "image/png": true}
    if _, allowed := allowedMIMEs[header.Header.Get("Content-Type")]; !allowed {
        errors.HandleError(w, errors.NewAPIError(nil, "Invalid file type", http.StatusBadRequest))
        return
    }

    // Create uploads directory if not exists.
    uploadsDir := "uploads"
    os.MkdirAll(uploadsDir, os.ModePerm)

    // Create a unique file name and save the file.
    fileName := "event_" + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(header.Filename)
    dstPath := filepath.Join(uploadsDir, fileName)
    dst, err := os.Create(dstPath)
    if err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Error creating file", http.StatusInternalServerError))
        return
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Error saving file", http.StatusInternalServerError))
        return
    }

    // Set the image URL in the event model.
    event.ImageURL = dstPath

    // Create the event using the service.
    if err := services.CreateEvent(event); err != nil {
        errors.HandleError(w, errors.NewAPIError(err, "Error creating event", http.StatusInternalServerError))
        return
    }

    // Respond with a success message.
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct {
        Message string `json:"message"`
    }{"Event created successfully"})
}