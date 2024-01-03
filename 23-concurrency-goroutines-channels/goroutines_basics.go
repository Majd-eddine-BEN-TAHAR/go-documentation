// goroutines_basics.go

package main

import (
	"fmt"
	"time"
)

// sayHello is a function that demonstrates a basic goroutine.
func sayHello() {
    fmt.Println("Hello from goroutine")
}

// demonstrateBasicGoroutine shows how to start a goroutine and the importance of synchronization.
func demonstrateBasicGoroutine() {
    // Starting the goroutine using the 'go' keyword.
    go sayHello()

    // Wait for a moment (1 second in this case) to allow the goroutine to execute.
    // This time.Sleep is used to pause the main function briefly, ensuring that the goroutine has time to run.
    // Without this delay, the main function might launch and exit before the goroutine has a chance to print its message.
    // As a result, "fmt.Println("Hello from goroutine")" might not get a chance to execute.
    time.Sleep(1 * time.Second)
    fmt.Println("Main function ends")
}
