package main

import (
	"fmt"
	"math/rand"
	"time"
)

// performTask simulates a task that can either complete successfully or fail.
// It represents a typical asynchronous operation, such as a network request or a long computation.
// 'id' is used to identify the task, 'success' channel is used to signal successful completion,
// and 'fail' channel is used to signal an error.
func performTask(id int, success chan<- string, fail chan<- error) {
    duration := rand.Intn(3) + 1 // Simulating variable task duration.
    time.Sleep(time.Duration(duration) * time.Second) // Simulating task processing time.

    // Randomly determining if the task fails or succeeds.
    if rand.Float32() < 0.3 { // Assigning a 30% chance for the task to fail.
        // Sending an error message to the 'fail' channel if the task fails.
        fail <- fmt.Errorf("task %d failed", id)
    } else {
        // Sending a success message to the 'success' channel if the task succeeds.
        success <- fmt.Sprintf("Task %d completed successfully in %d seconds", id, duration)
    }
}

// demonstrateSelectStatement showcases the use of the select statement in concurrent Go programming.
// It handles multiple asynchronous tasks, each of which can either complete successfully or fail.
// The function waits for all tasks to complete, handling their results as they come.
func demonstrateSelectStatement() {
    success := make(chan string) // Channel for successful task completions.
    fail := make(chan error)     // Channel for task failures.
    numberOfTasks := 5           // Total number of tasks to handle.

    // Starting each task in its own goroutine.
    // This allows tasks to run concurrently, independent of each other.
    for i := 1; i <= numberOfTasks; i++ {
        go performTask(i, success, fail)
    }

    completedTasks := 0 // Counter for completed tasks.

    // Continuously listening for results from the tasks.
    // The loop runs until all tasks are accounted for (either success or failure).
    for completedTasks < numberOfTasks {
        select {
        case result := <-success:
            // Handling a successful task completion.
            fmt.Println(result)
            completedTasks++ // Incrementing the counter for completed tasks.
        case err := <-fail:
            // Handling a task failure.
            fmt.Printf("Error: %s\n", err)
            completedTasks++ // Incrementing the counter for completed tasks.
        }
    }
}

