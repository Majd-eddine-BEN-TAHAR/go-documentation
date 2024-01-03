package main

import (
	"fmt"
)

// demonstrateChannels shows the basic creation and use of channels in Go.
func demonstrateChannels() {
    // Creating a new channel of integers.
    ch := make(chan int)

    // Starting a new goroutine.
    go func() {
        // Sending a value into the channel.
        ch <- 42
    }()

    // Receiving the value sent on the channel.
    // This line will block tha main func until a value is received from the channel.
    result := <-ch
    fmt.Println("Received:", result)
}
