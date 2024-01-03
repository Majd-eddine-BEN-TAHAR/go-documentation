package main

import "fmt"

// ----- Basic Function Definitions and Usage -----

// myFunction: Simple function demonstrating a function with no parameters and no return value.
// Purpose: To print a basic message to the console.
func myFunction() {
    // Print a greeting message to the console.
    fmt.Println("This is a function")
}

// functionWithoutReturn: Function that takes a string parameter (name) but does not return any value.
// Purpose: To greet a user by their name.
func functionWithoutReturn(name string) {
    // Output a personalized greeting message.
    fmt.Println("Hello", name)
}

// add: Function that takes two integers and returns their sum.
// Returns: Integer sum of 'a' and 'b'.
func add(a int, b int) int {
    return a + b
}

// subtract: Function that takes two integers and returns their difference.
// Returns: Integer difference between 'a' and 'b'.
func subtract(a int, b int) int {
    return a - b
}

// swap: Function returning multiple values.
// Takes two strings and swaps them.
// Returns: Two strings, reversed in order.
func swap(x, y string) (string, string) {
    return y, x
}

// divideAndMultiply: Performs division and multiplication on two integers.
// Returns: The result of division and multiplication
func divideAndMultiply(a, b int) (int, int) {
    divisionResult := a / b
    multiplicationResult := a * b
    return divisionResult, multiplicationResult
}

// example of Variadic function
// a Variadic function lets you pass a different number of arguments to a function. This is helpful when you are not sure how many arguments you might need to use.
// Sum takes any number of integer arguments and returns their sum
// (...int) means it accepts any number of int arguments
// the 3 dots means it accepts any number of int arguments
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

// ----- Advanced Function Concepts -----

// transform: Higher-order function that applies a transformation function to an integer.
// Parameters: 'num' - integer to transform, 'fn' - function defining the transformation.
// Returns: Transformed integer result.
func transform(num int, fn func(int) int) int {
    return fn(num)
}

// double: Example transformation function to double the input.
func double(num int) int {
    return num * 2
}

// triple: Example transformation function to triple the input.
func triple(num int) int {
    return num * 3
}

// ----- Using Type Aliases with Functions -----

// operationFunc: Type alias for functions that take two integers and return an integer.
type operationFunc func(int, int) int

// operate: Demonstrates the use of the operationFunc type alias.
// Parameters: 'a' and 'b' - integers, 'op' - function of type operationFunc.
// Returns: Result of applying 'op' to 'a' and 'b'.
func operate(a int, b int, op operationFunc) int {
    return op(a, b)
}

// ----- Main Function: Demonstrations -----

func main() {
    // simple function with no parameters and no return value.
    myFunction()
    
    // a function with a parameter but no return value.
    // Function `functionWithoutReturn` is called with "Alice" as an argument.
    functionWithoutReturn("Alice")

    
    // a function that returns a single value.
    fmt.Println("Sum:", add(5, 3))
    
    // a function that returns multiple values.
    a, b := swap("hello", "world")
    fmt.Println("Swapped:", a, b)
    
    // a function that performs multiple operations and returns multiple results.
    // Function `divideAndMultiply` is called with 10 and 5, and its results are printed.
    divResult, mulResult := divideAndMultiply(10, 5)
    fmt.Println("Division:", divResult, "Multiplication:", mulResult)

    // use of higher-order functions with a simple transformation.
    // Function `transform` is called with 5 and `double`, and its result is printed.
    // It's crucial to pass the function reference (e.g., `double`) without parentheses.
    // If parentheses were added (e.g., `double()`), it would call the function instead of passing it.
    fmt.Println("Doubled:", transform(5, double))
    
    // use of higher-order functions with a different transformation.
    // Function `transform` is called with 5 and `triple`, and its result is printed.
    // It's crucial to pass the function reference (e.g., `triple`) without parentheses.
    // If parentheses were added (e.g., `triple()`), it would call the function instead of passing it.
    fmt.Println("Tripled:", transform(5, triple))

    // Using an anonymous function as an argument.
    // anonymous functions doesn't have a name
    // The anonymous function doubles its input.
    result := transform(10, func(num int) int {
        return num * 2
    })
    fmt.Println("Result from anonymous function:", result)

    
    // Demonstrate the use of type aliases in functions, the type alias is used on the operate function parameter
    // Function `operate` is called with 10, 5, and `add`, and its result is printed.
    // It's crucial to pass the function reference (e.g., `add`) without parentheses.
    // If parentheses were added (e.g., `add()`), it would call the function instead of passing it.
    fmt.Println("Sum using operate:", operate(10, 5, add))
    
    // Demonstrate the use of type aliases with a different operation, the type alias is used on the operate function parameter
    // Function `operate` is called with 10, 5, and `subtract`, and its result is printed.
    fmt.Println("Difference using Alias:", operate(10, 5, subtract))

    // call the Variadic Sum with any number of arguments
    fmt.Println(sum(1, 2))       // Sum of 2 numbers
    fmt.Println(sum(1, 2, 3, 4)) // Sum of 4 numbers
    fmt.Println(sum())           // Sum of no numbers, returns 0
}