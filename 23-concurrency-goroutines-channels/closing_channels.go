package main

import (
	"fmt"
	"time"
)

// sendData simulates sending data over a channel.
// It sends a series of integers, then closes the channel to signal that no more data will be sent.
func sendData(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i // Sending data to the channel.
        time.Sleep(time.Second) // Simulating processing time.
    }
    close(ch) // Important: closing the channel to indicate that no more data will be sent.
}

// demonstrateClosingChannels shows the best practices for closing channels.
// It demonstrates how to send data to a channel, close it, and safely receive data until the channel is closed.
func demonstrateClosingChannels() {
    ch := make(chan int)

    go sendData(ch) // Start goroutine to send data.

    // Receiving data from the channel using a range loop.
    // The loop automatically breaks when the channel is closed, preventing deadlocks errors.
    // the main program will be blocked until we close the channel
    for value := range ch {
        fmt.Println("Received:", value)
    }

    fmt.Println("Channel closed and all data received")
}

