package main

import (
	"fmt"
	"os"
)

func main() {
	// Example of Scan
	var name string
	var age int
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age) // Reads and stores it in name and age, note:scan only reads one word
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Example of Sprintf
	// Sprintf is used for string formatting
	formattedString := fmt.Sprintf("Formatted: Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)

	// Example of Fprintf
	// Fprintf is used to write formatted strings to a writer (e.g., file)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "Name: %s, Age: %d\n", name, age) // Writes formatted string to file

	// Example of multi-line strings using backticks
	// Useful for having strings that span multiple lines
	multiLineString := `This is a multi-line string.
It can span multiple lines
without needing explicit newline characters.`
	fmt.Println(multiLineString)
}