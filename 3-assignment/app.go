package main

import "fmt"

// Go Variable Declaration and Reassignment Documentation
// first of all := and = are not the same, = is for assiging a value and := is for declaring a variable
func main() {
    // Using := (Short Variable Declaration)
    // - := is a concise way to declare and initialize a variable in one line.
    // - The type of the variable is inferred from the value it was initialized with.
    // - Once a variable is declared with a certain type using :=, its type cannot be changed within the same scope.
    // - Reassignment must be made with the inferred variable's initial type.

    // Example:
    age := 25 // 'age' is inferred as int
    // age = "thirty" // Compile-time Error: cannot assign a string to an int variable
	fmt.Println(age)


	// Using var for Explicit Variable Declaration
    // - var is used for more traditional, explicit variable declaration.
    // - It requires specifying the type of the variable.
    // - A variable declared with var can be reassigned, but only with values of the same type.
    // - If a different type of value is needed, it must be converted to the declared type of the variable

    // Example:
    var price int
    price = 20 // Assigning an int value
    floatPrice := 25.5 // inferred type as float64
    price = int(floatPrice) // Converting float64 to int and reassigning
	fmt.Println(price)

    // Attempting to assign a different type without conversion will result in an error:
    var message string
    message = "Price is 20"
    // message = 30 // Compile-time Error: cannot directly assign an int to a string variable
	fmt.Println(message)

    // Summary:
    // - Use := for quick and type-inferred declarations within functions.
    // - Use var for explicit type declarations, especially when the type is crucial, or for package-level variables.
    // - Reassignment must respect the variable's original type.
    // - To assign a value of a different type, conversion is required if it is permissible in Go.
}