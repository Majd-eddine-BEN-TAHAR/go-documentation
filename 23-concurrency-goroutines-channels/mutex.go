package main

import (
	"fmt"
	"os"
	"sync"
)

func demonstrateMutexUsage() {
    // 'os.O_CREATE' means the file will be created if it doesn't exist.
    // 'os.O_WRONLY' opens the file in write-only mode.
    // Combined, these flags ensure the file is created if not present and is opened for writing.
    file, err := os.OpenFile("shared_file.txt", os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var wg sync.WaitGroup
    var mutex sync.Mutex

    // Function for writing to the file
    writeToFile := func(id int) {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            mutex.Lock() // Lock the mutex to ensure exclusive access to the file.
            _, err := file.WriteString(fmt.Sprintf("Goroutine %d, iteration %d\n", id, i))
            mutex.Unlock() // Unlock the mutex to allow other goroutines to write.

            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
        }
    }

    
    // Start four goroutines
	wg.Add(4)
	for i := 1; i <= 4; i++ {
		go writeToFile(i)
	}

    // Wait for both goroutines to finish
    wg.Wait()
    fmt.Println("Finished writing to file")
}