package main

import (
	"fmt"
)

func main() {
    /* 
        Type Assertions with 'any' Type in Go
        - Type assertion(التأكد من النوع) in Go is a mechanism used to check if an interface value holds a specific type. This is particularly useful when you have a value of type interface{} and you need to determine its actual type and work with it in a type-specific way.
        - 'any' is an alias for the 'interface{}' type, introduced to improve code readability.
        - 'interface{}' is a special type that can hold values of any other type, known as an empty interface.
        - Type assertion is used to retrieve the dynamic value from a variable of type 'any'.
        - It allows checking the type of an 'any' variable and accessing its underlying value.
        - Type assertions can be used for type checking and extraction in a single operation.
    */

    // Example with Type Assertions:
    var value any = "Go is fun"
    // To extract the underlying value or check its type, a Type Assertion is used.

    // Performing a Type Assertion:
    str, ok := value.(string)
    // The assertion returns the value and a boolean indicating success.

    if ok {
        fmt.Printf("String value: %s\n", str)
    } else {
        fmt.Println("Value is not a string")
    }

    // Using Type Assertions in a switch statement:
    switch v := value.(type) {
    case string:
        fmt.Printf("String value: %s\n", v)
    case int:
        fmt.Printf("Integer value: %d\n", v)
    default:
        fmt.Println("Unknown type")
    }

    // Note: Type assertions can cause a panic if the assertion is wrong and not handled properly.
}