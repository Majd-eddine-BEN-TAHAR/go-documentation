package main

import (
	"fmt"
)

// intSeq returns a closure.
// Each closure maintains its own state with a unique 'i'.
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    // --- Example 1: Basic Closure ---
    // A basic closure: A function that is returned from another function.
    // The 'intSeq' function returns another function, which we define anonymously in its body.
    // 'intSeq' is a simple closure that increments a counter.
    nextInt := intSeq()

    // The closure captures its own 'i' value.
    fmt.Println(nextInt()) // Output: 1
    fmt.Println(nextInt()) // Output: 2
    fmt.Println(nextInt()) // Output: 3

    // To confirm that the state is unique to that particular function,
    // create and test a new one.
    newInts := intSeq()
    fmt.Println(newInts()) // Output: 1

    // --- Example 2: Closure Capturing External Variable ---
    // Closures can capture and modify variables defined outside the function body.
    start := 10
    incrementer := func() int {
        start++
        return start
    }

    fmt.Println(incrementer()) // Output: 11
    fmt.Println(incrementer()) // Output: 12

    // --- Example 3: Common Pitfall(misunderstood) with Loop Variable Capture ---
    // Explanation:
    // - Each closure in `funcs` refers to the same `i`.
    // - By the time closures are executed, the loop has finished, and `i` is 3.
    // - Common pitfall: assuming each closure captures the current value of `i`.
    var funcs []func()
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i)  // Incorrectly prints 3, 3, 3 if not handled correctly.
        })
    }

    for _, f := range funcs {
        f()  // Incorrectly prints 3, 3, 3
    }


    // --- Example 4: Correcting the Pitfall with Loop Variable Capture ---
    // Explanation:
    // - By re-declaring `i` in each iteration, each closure gets its own `i`.
    // - This solves the common pitfall of loop variable capture.
    funcs = nil // Resetting the slice
    for i := 0; i < 3; i++ {
        i := i // create a new `i` for each iteration
        funcs = append(funcs, func() {
            fmt.Println(i)
        })
    }

	
    for _, f := range funcs {
        f()  // Correctly prints 0, 1, 2
    }
}