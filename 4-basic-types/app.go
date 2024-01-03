package main

import (
	"fmt"
)

func main() {
    /*
		Go is a statically typed language. This means that variables have a specific type, and that type cannot change.
    */

    // Basic Types in Go:
    // 1. Integers: These are whole numbers without a decimal. Examples include int, int8, int16, int32, int64.
    var myInt int = 10
    fmt.Println("Integer example:", myInt)

    // 2. Floating Point Numbers: These are numbers with decimals. Examples include float32, float64.
    var myFloat float64 = 3.14
    fmt.Println("Floating point example:", myFloat)

    // 3. Strings: A string in Go is a sequence of characters. Strings are immutable.
	// When you reassign the myString variable with a new string value "majd," it doesn't modify the original string. Instead, it creates a completely new memory allocation for the string "majd," and the myString variable is updated to point to this new memory location.
    var myString string = "Hello, Go!"
	myString = "majd"
    fmt.Println("String example:", myString)

    // 4. Booleans: This type represents a truth value, which can be either true or false.
    var myBool bool = true
    fmt.Println("Boolean example:", myBool)

	// -----------------------------------------------------------------------------
    /*  
		4.Arrays:
	*/
	// An array is a fixed-size sequence of elements of the same type. A key property of an array is its size, which is part of its type. Once you define an array with a certain size, it cannot be resized, distinguishing it from slices which are more flexible.
	// You can only modify the values of individual elements within the array using standard indexing and assignment operations.
          
    // Limitations of Arrays:
	// Arrays have a fixed size. If you need a dynamic size, consider using slices.
	// Slices can grow or shrink and are more commonly used than arrays for this flexibility.

    // Declaring and Initializing an Array:
	// Here, 'myArray' is an array of 3 integers. This array is initialized with values {1, 2, 3}.
	// The size [3]int is part of the array's type and cannot be changed.
	var myArray [3]int = [3]int{1, 2, 3}

	// Modifying an Array:
	// Arrays are zero-indexed in Go, meaning the first element is at index 0.
	// Here, we change the first element (index 0) from 1 to 5.
	myArray[0] = 5

	// Accessing Array Elements:
	// You can access elements of an array using an index. Here, we print the modified array.
	// Output will be [5 2 3], since only the first element was changed.
	fmt.Println("Array example:", myArray)

	// Iterating over an Array:
	// You can iterate over an array using a 'for' loop. The 'range' keyword provides both
	// the index and the value at that index.
	fmt.Println("Iterating over the array:")
	for index, value := range myArray {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	

}