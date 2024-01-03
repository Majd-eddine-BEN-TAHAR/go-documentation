package main

import (
	"fmt"
	"time"
)

// User represents a user in the system.
type User struct {
    Name  string
    Email string
}

func unbufferedChannelExample() {
    newUser := User{Name: "John Doe", Email: "john@example.com"}
    done := make(chan bool) // Unbuffered channel for synchronization.

    go processRegistration(newUser, done)

    // Wait for the registration to complete.
    // This will block the main program from execution until a message is received from the 'done' channel.
    <-done

    // Continue execution after registration is done.
    sendWelcomeEmail(newUser)
}

// processRegistration handles user registration logic.
// In a real scenario, this would add the user to the database.
func processRegistration(user User, done chan bool) {
    // Simulating some work to make it appear that it's really happening.
    time.Sleep(2 * time.Second)

    // Notify that the registration is done.
    done <- true
}

// sendWelcomeEmail sends a welcome email to the user.
func sendWelcomeEmail(user User) {
    fmt.Printf("Sending welcome email to %s\n", user.Email)
}