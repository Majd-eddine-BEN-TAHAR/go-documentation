package main

import (
	"fmt"
)

func main() {
    // Example 1: Simple for loop
    // This loop starts with i = 0 and runs while i < 5, increasing i by 1 each time.
    for i := 0; i < 5; i++ {
        fmt.Println("Simple Loop, iteration:", i)
    }

    // Example 2: For loop with a condition only (while-like loop)
    // This loop runs while j is less than 3, starting with j = 0.
    j := 0
    for j < 3 {
        fmt.Println("Condition-only Loop, iteration:", j)
        j++ // Increase j by 1
    }

    // Example 3: Infinite loop with a break condition
    // This loop runs indefinitely until the break condition is met.
    k := 0
    for {
        fmt.Println("Infinite Loop, iteration:", k)
        k++
        if k == 2 { // When k equals 2, exit the loop
            break
        }
    }

    // Example 4: For loop with continue
    // This loop skips the current iteration when l is even.
    for l := 0; l < 5; l++ {
        if l%2 == 0 { // Check if l is even
            continue // Skip the rest of the loop for this iteration
        }
        fmt.Println("Loop with Continue, iteration:", l)
    }

    // Example 5: For-each range loop
    // This loop iterates over each element of the slice.
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Println("Range Loop, index:", index, "value:", value)
    }

	// Example 6: For-each range loop
    // This loop iterates over each element of the map.
	data := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
    }

    // Iterating over the map using a for-range loop
    for key, value := range data {
        fmt.Println("Key:", key, "Value:", value)
    }
}