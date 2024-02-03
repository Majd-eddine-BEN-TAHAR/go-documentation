package main

import "fmt"

func main() {
    fmt.Println(isBalanced("(()((())())(()))(()(()))"))          // true
    fmt.Println(isBalanced("(()((())())(())()))(()(()))"))       // false
    fmt.Println(isBalanced("(((()((())())(()))(()(()))))("))     // false
    fmt.Println(isBalanced("()()(()()()()))(()"))                // false
    fmt.Println(isBalanced("(()())())(()())"))                   // false
    fmt.Println(isBalanced("(()()())"))                          // true
    fmt.Println(isBalanced("(())((())())"))                      // true
    fmt.Println(isBalanced("(((((((((())))))))))"))              // true
    fmt.Println(isBalanced("(()(()((())()"))                     // false
    fmt.Println(isBalanced("((()(()))(()(()(()(())))))"))        // true
}

func isBalanced(str string) bool {
    // Early return for strings that start with ')' or end with '('
    if len(str) > 0 && (str[0] == ')' || str[len(str)-1] == '(') {
        return false
    }

    count := 0  // Initialize a counter for tracking the balance

    for _, char := range str {
        // Increment the counter for an opening parenthesis
        if char == '(' {
            count++
        } else if char == ')' {
            // Decrement the counter for a closing parenthesis
            count--

            // If the count goes negative, parentheses are imbalanced
            if count < 0 {
                return false
            }
        }
    }

    // If the counter is zero, parentheses are balanced
    return count == 0
}