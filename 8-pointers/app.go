package main

import "fmt"

// Function that takes two pointers to int, dereferences them, and returns their sum.
func add(a *int, b *int) int {
    return *a + *b
}

func main() {
    // Pointer in Go
    // In Go, each pointer is a seprate type that represents the memory address of a value.
    // Pointers are considered types in Go.
    // A pointer type is defined in relation to an existing data type, indicating it can hold the address of a value of that specific type.
    // For example, a '*int' pointer type can only hold the address of an 'int' variable.
    // This type-specific approach to pointers ensures type safety and aids in writing clear, maintainable code.
    // ---------------------

    // Declaring a Pointer:
    // A pointer is declared with the '*' prefix. It stores the memory address of a variable.
    var ptr *int
    fmt.Println("Initial value of pointer:", ptr) // Initially, the pointer is nil.

    // Assigning an Address to a Pointer:
    // The '&' operator is used to get the address of a variable.
    var number int = 42
    ptr = &number
    fmt.Println("Address stored in ptr:", ptr)

    // Dereferencing(Getting the actual stored value in) a Pointer:
    // The '*' prefix is used to dereference the value at the memory address pointed by the pointer.
    // This is called 'dereferencing' the pointer.
    fmt.Println("Value of the variable ptr points to:", *ptr)

    // Changing the Value through a Pointer:
    // You can change the value at the referenced address using the pointer.
    *ptr = *ptr + 1
    fmt.Println("New value of the variable ptr points to:", *ptr)
    fmt.Println("Value of number after modification:", number)

    // "--------------------------------------------------------------"

    // Pointer to Pointer:
    // A pointer can also store the address of another pointer, leading to multiple levels of indirection.
    var num int = 110

    var pointer1 *int = &num
    *pointer1 = *pointer1 + 1

    var pointer2 **int = &pointer1
    **pointer2 = **pointer2 + 1

    var pointer3 ***int = &pointer2
    ***pointer3 = ***pointer3 + 1

    fmt.Println(pointer1)
    fmt.Println(*pointer1)
    fmt.Println(pointer2)
    fmt.Println(**pointer2)
    fmt.Println(pointer3)
    fmt.Println(***pointer3)
    fmt.Println("----------")
    fmt.Println(num)

    // "--------------------------------------------------------------"
    // Dynamic Memory Allocation for Pointers:
    // In Go, the 'new' function is used to allocate memory for a specific data type at runtime.
    // It allocates space for a variable and returns a pointer to it.
    // This dynamic memory allocation feature is useful when you need to work with data structures like maps or slices
    // whose size or lifetime is not known at compile time.

    // Declare a pointer variable 'ptrDynamic' for an integer.
    // Allocate memory for an integer and assign the pointer to 'ptrDynamic'.
    ptrDynamic := new(int)

    // Print the memory address where the integer is stored.
    fmt.Println("Memory address of dynamically allocated integer:", ptrDynamic)

    // Assign a value of 30 to the dynamically allocated integer through the pointer.
    *ptrDynamic = 30

    // Print the value of the dynamically allocated integer.
    fmt.Println("Value of dynamically allocated integer:", *ptrDynamic)

    // "--------------------------------------------------------------"

    // Compare Pointers :
    // We can compare pointers using the == and != operators.
    // These comparisons check if pointers are nil or if they point to the same address.
    var number3 int = 20
    // Create a pointer to an integer and assign it the address of 'number'.
    var pointer4 *int = &number3

    // Create another pointer to an integer and assign it the address of 'number'.
    var pointer5 *int = &number3

    // Check if pointer4 and pointer5 point to the same address.
    fmt.Println("Are pointer4 and pointer5 pointing to the same address?", pointer4 == pointer5)


    // ----------------------------------------------------------------------------------------
    // ----------------------------------------------------------------------------------------
    // example of passing pointers to function and i will dereference them in the function definiton
    // Define two integers.
    x := 3
    y := 4

    // Call add with pointers to x and y.
    sum := add(&x, &y)

    // Print the result.
    fmt.Println("Sum:", sum)
}