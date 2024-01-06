package models

import "time"

type Booking struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    EventID   int       `json:"event_id"`
    Timestamp time.Time `json:"booking_time"`
}