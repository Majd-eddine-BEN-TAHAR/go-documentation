package main

import (
	"log"
	"net/http"

	"event_booking_api/internal/app"
	"event_booking_api/internal/config"
	"event_booking_api/pkg/database"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()
    // Initialize the database and create it if it doesn't exist
    if err := database.InitDB(cfg.DatabasePath); err != nil {
        log.Fatalf("Could not set up database: %v", err)
    }
    defer database.DB.Close()

	// Initialize database schema and create tables
    database.InitDBSchema(database.DB, "./scripts")

    // Serve files from the 'uploads' directory under the '/uploads' URL path
    http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
    
    // main router 
    http.HandleFunc("/", app.Router)

    // Define the HTTP server
    httpPort := cfg.PORT
    if httpPort == "" {
        httpPort = "3000" // Default port if not specified
    }

    // Start the HTTP server
    log.Printf("Starting server on port %s", httpPort)
    if err := http.ListenAndServe(":"+httpPort, nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
