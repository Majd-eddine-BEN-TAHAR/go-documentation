package main

import (
	"fmt"
	"time"
)

// Producer function
func producer(buffer chan<- int) {
	for i := 0; i < 10; i++ {
		buffer <- i // Sending data to the buffer
		fmt.Println("Produced", i)
		time.Sleep(time.Millisecond * 500) // Added a small delay for visibility
	}
	close(buffer) // Close the channel once done producing
}

// Consumer function
func consumer(buffer <-chan int, done chan<- bool) {
	for value := range buffer {
		fmt.Println("Consumed", value)
	}
	done <- true
}

// demonstrateProducerConsumer shows a basic producer-consumer scenario
func demonstrateProducerConsumer() {
	fmt.Println("\nProducer-Consumer Example:")
	sharedBuffer := make(chan int, 1) // Creating a buffered channel with capacity 1
	done := make(chan bool)

	go producer(sharedBuffer) // Starting the Producer goroutine
	go consumer(sharedBuffer, done) // Starting the Consumer goroutine

	// Wait for the consumer to finish processing
	<-done
	fmt.Println("Finished producing and consuming.")
}