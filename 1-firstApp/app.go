// Go Type Safety: Go is a statically typed language, which means that type errors are caught at compile-time rather than runtime

// Every Go file starts with a package declaration.
// Here, 'package main' indicates that this file belongs to the 'main' package.
// The 'main' package is special in Go; it's used for executables, not libraries.
package main

// Import statements are used to include code from other packages.
// This allows the file to use functions, types, variables, etc., from these packages.
// Below, we're importing the 'fmt' package, a part of the huge Go standard library.
// The Standard library includes a wide range of built-in packages, accessible to all Go programs.
// The 'fmt' package, in this context, serves as a fundamental tool, enabling essential formatting functionalities for input and output operations.
import "fmt"

// The 'main' function is the entry point of any executable Go program.
// When the program is run, the code inside the 'main' function is executed first.
// This function does not take any arguments and does not return any value.
// if you have 2 main functions in the same module you will see errors
func main() {
    // Here, we're using the 'Println' function from the 'fmt' package.
    // This function prints the supplied string followed by a newline to the standard output.
    // Since 'Println' starts with a capital letter, it is an exported name and can be used outside of the 'fmt' package.
    fmt.Println("Hello, World!")

	fmt.Print("Hello world!")
	fmt.Print(`Hello world!`)

	/* single quotes not allowed */
	// fmt.Print('Hello world!') 
}