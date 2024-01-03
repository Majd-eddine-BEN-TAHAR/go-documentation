package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// performHTTPRequests performs multiple HTTP GET requests concurrently using sync.WaitGroup.
func demonstratewaitGroup() {
    urls := []string{
        "https://jsonplaceholder.typicode.com/todos/1",
        "https://jsonplaceholder.typicode.com/todos/2",
        // Additional URLs can be added here.
    }

    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1) // Increment the WaitGroup counter for each URL.
        go func(url string) {
            defer wg.Done() // Decrement the counter upon completion of the goroutine.
            if resp, err := http.Get(url); err == nil {
                fmt.Println("Fetched", url, "Status:", resp.Status)
                body, _ := io.ReadAll(resp.Body)
                // Print the response body.
                fmt.Println(string(body))

                resp.Body.Close()
            } else {
                fmt.Println("Error fetching", url, ":", err)
            }
        }(url)
    }

    wg.Wait() // Wait for all goroutines to complete.
    fmt.Println("All requests completed")
}