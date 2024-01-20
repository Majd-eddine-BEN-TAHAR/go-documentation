package main

import "fmt"

func main() {
    // a basic goroutine
    demonstrateBasicGoroutine()

    // basic creation and use of channels
    demonstrateChannels()

    // demonstrateClosingChannels shows the best practices for closing channels.
    demonstrateClosingChannels()

    // buffered channel example
    fmt.Println("Buffered Channel Example:")
    bufferedChannelExample()

    // unbuffered channel example
    fmt.Println("\nUnbuffered Channel Example:")
    unbufferedChannelExample()

    // Demonstrate the Producer-Consumer pattern
    demonstrateProducerConsumer()

    // Running the demonstration of the select statement.
    demonstrateSelectStatement()

    // performs multiple HTTP GET requests concurrently using sync.WaitGroup
    demonstratewaitGroup()
}
