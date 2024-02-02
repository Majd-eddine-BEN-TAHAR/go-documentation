package main

import "fmt"

func main() {
    fmt.Println(findMissingNumber([]int{1, 2, 3, 5})) // Should return 4 (Missing number is 4)
	fmt.Println(findMissingNumber([]int{1, 2, 3, 4, 6})) // Should return 5 (Missing number is 5)
	fmt.Println(findMissingNumber([]int{1, 3, 4, 5, 6})) // Should return 2 (Missing number is 2)
}

// findMissingNumber finds the missing number in an array of numbers from 1 to n (inclusive).
// It returns the missing number.
func findMissingNumber(arr []int) int {
    arraySum := 0
    expectedSum := 0
    
    // Calculate the sum of elements in the given array.
    for i := 0; i < len(arr); i++ {
        arraySum += arr[i]
    }
    
    // Calculate the expected sum of numbers from 1 to n, where n is the largest number in the array.
	n := arr[len(arr)-1]
	for i := 1; i <= n; i++ {
    	expectedSum += i
	}
    
    // Calculate and return the missing number by subtracting the array sum from the expected sum.
    return expectedSum - arraySum
}