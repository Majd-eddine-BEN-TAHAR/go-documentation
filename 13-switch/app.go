package main

import (
	"fmt"
)

// ProcessValue takes an 'any' type and uses a switch to handle different types.
func ProcessValue(value any) {
    switch value := value.(type) {
    case string:
        fmt.Printf("String value: %s\n", value)
    case int:
        fmt.Printf("Integer value: %d\n", value)
    case float64:
        fmt.Printf("Float value: %f\n", value)
    case bool:
        fmt.Printf("Boolean value: %t\n", value)
    default:
        fmt.Printf("Unknown type: %T\n", value)
    }
}

func main() {
    // Example 1: Basic switch statement
    // A switch statement allows you to select one of several code blocks to be executed.
    // It evaluates the value of an expression and matches it with one of the case values.
    // When a match is found, the corresponding block of code is executed.
    day := "Wednesday"
    switch day {
    case "Monday":
        fmt.Println("It's Monday.")
    case "Tuesday":
        fmt.Println("It's Tuesday.")
    case "Wednesday":
        fmt.Println("It's Wednesday.")
    case "Thursday":
        fmt.Println("It's Thursday.")
    case "Friday":
        fmt.Println("It's Friday.")
    default:
        fmt.Println("It's the weekend.")
    }

    // Example 2: Switch with expressions
    // In addition to simple values, you can use expressions in switch cases.
    // The first case that evaluates to true will execute.
    number := 42
    switch {
    case number < 0:
        fmt.Println("Negative number")
    case number >= 0 && number < 10:
        fmt.Println("Single-digit number")
    case number >= 10 && number < 100:
        fmt.Println("Double-digit number")
    default:
        fmt.Println("Huge number")
    }

    // Example 3: Switch with a value of type 'any', check the function
    // This function, `ProcessValue`, takes an 'any' type and uses a switch
    // to perform different operations based on the actual type of the argument.
    ProcessValue("Hello, Go!")
    ProcessValue(100)
    ProcessValue(99.99)
    ProcessValue(true)
}