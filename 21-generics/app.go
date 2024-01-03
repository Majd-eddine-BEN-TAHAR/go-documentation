package main

import (
	"fmt"
)

// Introduction to Generics in Go
// --------------------------------
// Generics allow the creation of flexible, type-safe functions and data structures in Go.
// They were introduced to provide more reusability and reduce the need for type assertions and interface{}. Type assertion(التأكد من النوع) in Go is a mechanism used to check if an interface value holds a specific type. This is particularly useful when you have a value of type interface{} and you need to determine its actual type and work with it in a type-specific way.
// Generics enable writing functions and data types that work for multiple types without sacrificing type safety.
//
// Advantages:
// - Type Safety: Prevents runtime type errors.
// - Reusability: Write code that works with different types.
// - Readability: Clearer intent without sacrificing performance.
//
// Disadvantages:
// - Complexity: Can make code more complex or harder to read if used too much.
// - Compilation Time: May increase the compilation time.
//
// Syntax of Generics:
// - To write a generic function, put a list of types inside square brackets before the function's parameters. For example, func greet[T Type](args).
// - This list of types tells you what kinds of types you can use in the function.
// - You can set rules (using interfaces) on these types to control what you can do with them.

// operate is a generic function that performs an operation based on the type of its arguments.
// It supports addition for numeric types (int, float64) and concatenation for strings.
// The function utilizes Go's generics syntax, allowing it to work with multiple types.
func operate[T int | float64 | string](a, b T) T {
	// Depending on the type of T, it either adds (for numbers) or concatenates (for strings).
	return a + b
}

// operate2 demonstrates the use of interface{} for a similar operation but without generics.
// It tries to assert(التأكد) the type of its arguments and performs the corresponding operation.
// This method is less type-safe and can be more error-prone compared to generics.
func operate2(a, b interface{}) interface{} {
	// Try to assert the type of a and b to int.
	// If both are ints, perform addition.
	intA, okA := a.(int)
	intB, okB := b.(int)
	if okA && okB {
		return intA + intB
	}

	// Try to assert the type of a and b to float64.
	// If both are float64s, perform addition.
	floatA, okA := a.(float64)
	floatB, okB := b.(float64)
	if okA && okB {
		return floatA + floatB
	}

	// Try to assert the type of a and b to string.
	// If both are strings, perform concatenation.
	stringA, okA := a.(string)
	stringB, okB := b.(string)
	if okA && okB {
		return stringA + stringB
	}

	// If none of the assertions succeed, return nil.
	return nil
}

func main() {
	// Demonstrate the use of operate with different types.
	fmt.Println(operate(2, 3))             // Outputs: 5
	fmt.Println(operate(1.5, 2.3))         // Outputs: 3.8
	fmt.Println(operate("Hello, ", "Go!")) // Outputs: "Hello, Go!"

	// Demonstrate the use of operate2 with different types and a failed type assertion.
	fmt.Println(operate2(2, 3))             // Outputs: 5
	fmt.Println(operate2(1.5, 2.3))         // Outputs: 3.8
	fmt.Println(operate2("Hello, ", "Go!")) // Outputs: "Hello, Go!"
	fmt.Println(operate2(2, "3"))           // Outputs: nil (failed type assertion)
}