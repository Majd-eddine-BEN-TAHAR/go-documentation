package main

import (
	"fmt"
	"math"
)

func main() {
    // Test cases
    fmt.Println(findLargestNumber([]int{3, 9, 5, -1, 12, 0})) // Should return 12
    fmt.Println(findLargestNumber([]int{-10, -5, -20, -3}))   // Should return -3
    fmt.Println(findLargestNumber([]int{7}))                 // Should return 7
    fmt.Println(findLargestNumber([]int{}))                  // Should return math.MinInt32 (minimum int value)
}

func findLargestNumber(numbers []int) int {
    if len(numbers) == 0 {
        // Handle the case of an empty array by returning the minimum integer value
        return math.MinInt32
    }

    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}