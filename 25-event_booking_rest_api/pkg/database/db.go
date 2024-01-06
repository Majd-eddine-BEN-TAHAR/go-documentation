package database

import (
	"database/sql"
	"os"

	// In Go, when an underscore _ is used before an import statement, it indicates that the package is being imported for its side-effects only, and you won't directly use any functions, types, or variables from that package. This is known as a "blank import."
	// For github.com/mattn/go-sqlite3, the underscore is used because this package registers itself with the database/sql package when imported. The driver (go-sqlite3) is necessary for the sql package to work with SQLite databases, but you don't directly call any functions from go-sqlite3.
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// DB is a global variable to hold the database connection
var DB *sql.DB

// InitDB initializes the database connection using a given dataSourceName
func InitDB(dataSourceName string) error {
    // Check if the database file exists
    if _, err := os.Stat(dataSourceName); os.IsNotExist(err) {
        file, err := os.Create(dataSourceName) // Create the file if it does not exist
        if err != nil {
            return err
        }
        file.Close() // Close the file after creating it
    }
    
    var err error
    // Open a database connection with the SQLite driver and the provided dataSourceName
    DB, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return err
    }
    // Ping the database to verify the connection is established
    if err = DB.Ping(); err != nil {
        return err
    }
    return nil
}
