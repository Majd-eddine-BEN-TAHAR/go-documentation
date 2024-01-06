package models

import "time"

// Event represents the structure for an event.
type Event struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Location    string    `json:"location"`
    StartTime   time.Time `json:"start_time"`
    EndTime     time.Time `json:"end_time"`
    CreatorID int         `json:"creator_id"` // Reference to the user who created the event.
    ImageURL string       `json:"image_url,omitempty"` // URL of the event image
}
