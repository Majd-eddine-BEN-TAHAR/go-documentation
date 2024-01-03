// Understanding Zero Values in Go
// When you create a variable in Go without giving it a value, Go automatically gives it a default value. This is known as the variable's "zero value.”

package main

import "fmt"

func main() {

    // Integer Zero Value:
    // In Go, when an integer variable is declared but not initialized, its zero value is automatically set to 0.
    var age int
    // Here, 'age' is an integer with a zero value of 0.
    fmt.Println("Age:", age) // Output: 0

    // Float Zero Value:
    // Similarly, for float types, the zero value is 0.0.
    var temperature float64
    // The 'temperature' variable has a zero value of 0.0.
    fmt.Println("Temperature:", temperature) // Output: 0.0

    // String Zero Value:
    // For strings, the zero value is an empty string, denoted by "".
    var name string
    // 'name' is a string with a zero value of an empty string.
    fmt.Println("Name:", name) // Output: ""

    // Boolean Zero Value:
    // Boolean types in Go have 'false' as their zero value.
    var isActive bool
    // The boolean variable 'isActive' has a zero value of false.
    fmt.Println("Is Active:", isActive) // Output: false

    // Pointer Zero Value:
    // Pointers have a zero value of 'nil', indicating they don't point to any allocated memory.
    var pointer *int
    // 'pointer' is a pointer to an int, with a zero value of nil.
    fmt.Println("Pointer:", pointer) // Output: <nil>

	// Slice Zero Value:
	// A nil slice is a slice that doesn't reference any underlying array. It's not the same as an empty slice
    // A 'nil' slice in Go is displayed as [] when printed, not explicitly as 'nil'.
    // This output shows an empty slice, but it does not differentiate between 'nil' and an empty slice.
    var numbers []int
    // 'numbers' is a slice with a zero value of nil.
    fmt.Println("Numbers:", numbers) // Output: []

    // Map Zero Value:
	// it's important to differentiate between a nil map and an empty map; they are not equal.
    // A 'nil' map is displayed as map[] when printed, similar to slices.
    // This output indicates an empty map, without differentiating between 'nil' and an empty initialized map.
    var user map[string]string
    // The map 'user' has a zero value of nil.
    fmt.Println("User:", user) // Output: map[]

    // Channel Zero Value:
    // Channels have a zero value of 'nil'. A 'nil' channel is not usable for sending or receiving.
    var messages chan string
    // 'messages' is a channel with a zero value of nil.
    fmt.Println("Messages:", messages) // Output: <nil>
}

// Key Points:
// - Zero values in Go provide default states for variables, ensuring they are initialized.
// - Understanding zero values is crucial for robust programming and preventing nil pointer dereferences.
// - Go's approach to zero values, especially with 'nil', is distinct from 'null' in many other languages.
// - Go's fmt package prints zero values for slices and maps as empty structures, not as 'nil'.