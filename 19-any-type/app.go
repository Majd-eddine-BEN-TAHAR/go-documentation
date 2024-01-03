package main

import "fmt"

// PrintInfo prints the value and its type, accepting any type of value.
func PrintInfo(v any) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {

    /*  
        The 'any' type (previously 'interface{}' before go 1.18)
            - The 'any' type is a way to represent any type in Go. It's useful when the exact type is unknown or can vary.
            - 'any' is an alias for 'interface{}', introduced in Go 1.18 for better readability and understanding.
    */

    // Example of 'any' type:
    var value any = "Hello, Go!" // 'value' can be of any type, here it's a string
    fmt.Println("Value:", value)

    // Using 'any' in a slice:
    var mixed []any = []any{"Go", 42, true} // A slice that can hold values of any type
    fmt.Println("Mixed Slice:", mixed)

    // Type assertion with 'any':
    // To use the value in its specific type, we perform a type assertion.
    // This checks if 'value' is of a certain type (e.g., string) and converts it.
    str, ok := value.(string) // Converts 'value' to 'string' type if possible
    if ok {
        fmt.Println("String Value:", str)
    } else {
        fmt.Println("Value is not a string")
    }

    // Using 'any' in a map:
    var settings map[string]any = make(map[string]any) // A map with string keys and values of any type
    settings["version"] = 1.0
    settings["name"] = "GoLang Project"
    settings["debug"] = true
    fmt.Println("Settings:", settings)

    // Using 'any' allows for flexible function parameters:
    // A function that can accept any type of argument.
    PrintInfo("String")
    PrintInfo(123)
    PrintInfo(false)
}

