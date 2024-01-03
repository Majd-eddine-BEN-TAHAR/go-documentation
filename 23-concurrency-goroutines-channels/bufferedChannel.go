package main

import "fmt"

func bufferedChannelExample() {
    // Creating a buffered channel of integers with a capacity of 3
    bufferedChan := make(chan int, 3)

    // Sending values into the channel without blocking
    bufferedChan <- 1
    bufferedChan <- 2
    bufferedChan <- 3

    // Receiving values from the channel
    for i := 0; i < 3; i++ {
        value := <-bufferedChan
        fmt.Println("Received value:", value)
    }
}

