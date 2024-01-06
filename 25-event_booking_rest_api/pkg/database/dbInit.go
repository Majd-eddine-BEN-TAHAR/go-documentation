package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

// InitDBSchema initializes the database schema using SQL scripts
func InitDBSchema(db *sql.DB, scriptsPath string) {
    // Read SQL file
    path := filepath.Join(scriptsPath, "init_db.sql")
    content, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("Error reading SQL file: %v", err)
    }

    // Execute SQL commands
    _, err = db.Exec(string(content))
    if err != nil {
        log.Fatalf("Error executing SQL script: %v", err)
    }

    log.Println("Database schema initialized successfully")
}