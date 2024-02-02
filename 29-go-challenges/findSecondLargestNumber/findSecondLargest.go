// finds the second largest number in an integer slice.
package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(findSecondLargestNumber([]int{3, 9, 5, -1, 12, 0})) // Should return 9
	fmt.Println(findSecondLargestNumber([]int{-10, -5, -20, -3}))  // Should return -5
	fmt.Println(findSecondLargestNumber([]int{7}))                // Should return 0 (as there is no second largest number)
	fmt.Println(findSecondLargestNumber([]int{5, 5, 5}))           // Should return 5 (as duplicates are allowed)
	fmt.Println(findSecondLargestNumber([]int{}))                 // Should return 0 (as the array is empty)
}

func findSecondLargestNumber(arr []int) int {
	largest := math.MinInt32
	secondlargest := math.MinInt32
	for _, num := range arr {
		// If the current element 'num' is greater than 'secondlargest', update 'secondlargest' to 'num'.
		if(num > secondlargest){
			secondlargest = num
		}
		// If 'secondlargest' is now greater than 'largest', swap the values of 'largest' and 'secondlargest'.
		if(secondlargest > largest){
			temp := largest
			largest = secondlargest
			secondlargest = temp
		}
	}
	// After the loop, if 'secondlargest' is still the minimum integer value (math.MinInt32), set it to 0.
	if secondlargest == math.MinInt32{
		secondlargest = 0
	}
	return secondlargest
}