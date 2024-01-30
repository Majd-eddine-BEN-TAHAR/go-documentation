package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
    // Test cases
    fmt.Println(isPalindrome("racecar"))              			  // true
    fmt.Println(isPalindrome("A man, a plan, a canal, Panama!"))  // true
    fmt.Println(isPalindrome("hello"))                 			  // false
    fmt.Println(isPalindrome("Was it a car or a cat I saw?"))     // true
    fmt.Println(isPalindrome("12321"))                 			  // true
    fmt.Println(isPalindrome("No 'x' in Nixon"))       			  // true
    fmt.Println(isPalindrome("This is not a palindrome")) 		  // false
    fmt.Println(isPalindrome(""))                      			  // true (an empty string is considered a palindrome)
}

// isPalindrome checks if a given string is a palindrome.
// It returns true if the input string is a palindrome and false otherwise.
func isPalindrome(str string) bool {
    // Remove non-alphabetic characters and convert the string to lowercase
    regexp := regexp.MustCompile("[^a-zA-Z]+")
    inputString := strings.ToLower(regexp.ReplaceAllString(str, ""))

    // Calculate the length of half the string
    halfLength := len(inputString) / 2

    // Iterate through the first half of the string and compare characters with their mirrored counterparts
    for i := 0; i < halfLength; i++ {
        if inputString[i] != inputString[len(inputString)-1-i] {
            // If characters do not match, it's not a palindrome
            return false
        }
    }

    // If the loop completes without finding a mismatch, it's a palindrome
    return true
}