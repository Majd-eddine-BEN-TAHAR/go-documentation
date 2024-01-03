package main

import (
	"fmt"

	// The 'using_our_packages/' directory is a module that contains various utility packages created for this project, we did that when we initialize  module using => go mod init using_our_packages , so it's the name of our modeule(project) and we should include it in the path to let go understand we want to import a local module
	// 'myMathUtils' is one of the packages inside this module, providing mathematical functions like Add and Subtract.
	// Import the myMathUtils package from a local module. This is a custom package i created, it's in the directory
	"using_our_packages/myMathUtils"

	// Import the govalidator package for data validation. This package is a third party package from github and it  provides a set of functions to perform validations on strings, numerics, slices, and structs.
	"github.com/asaskevich/govalidator"
)

func main() {
	a := 10
	b := 5

	// Use the Add function from the myMathUtils package to calculate the sum of a and b.
	// This demonstrates how to use a function from a custom package.
	sum := myMathUtils.Add(a, b)

	// Use the Subtract function from the myMathUtils package to calculate the difference between a and b.
	// This is another example of using a custom package function.
	difference := myMathUtils.Subtract(a, b)

	// Print the results using fmt.Printf. The %d verb is used to format integers.
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)

	// ---------------------------------------------------------
	// Using the govalidator package. This section demonstrates using an external package.
	// Note: The govalidator package must be installed before running this code using => go get github.com/asaskevich/govalidator
	// Email address to be validated.
	email := "example@example.com"

	// Use the IsEmail function from the govalidator package to check if the email string is a valid email format.
	isValid := govalidator.IsEmail(email)

	if isValid {
		fmt.Printf("%s is a valid email address.\n", email)
	} else {
		fmt.Printf("%s is not a valid email address.\n", email)
	}
}