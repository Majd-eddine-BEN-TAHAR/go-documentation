package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
    // Writing to a File
    // To write data to a file in Go, you can follow these steps:

    // Step 1: Create or open a file for writing.
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
	// In Go, defer is a keyword used to schedule a function call to be executed when the surrounding function (in this case, main) finishises executing. In the context of file handling, defer file.Close() is used to ensure that the file is closed properly when you're done with it, even if an error occurs or if there are multiple return statements within the function.
    defer file.Close()

    // Step 2: Prepare the data you want to write (in this case, a simple string).
    data := "Hello, World!\n"

    // Step 3: Write the data to the file.
    _, err = file.WriteString(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    fmt.Println("Data written to file successfully.")

    // Reading from a File
    // To read data from a file in Go, you can follow these steps:

    // Step 4: Open the file for reading.
    file, err = os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
	// In Go, defer is a keyword used to schedule a function call to be executed when the surrounding function (in this case, main) returns. In the context of file handling, defer file.Close() is used to ensure that the file is closed properly when you're done with it, even if an error occurs or if there are multiple return statements within the function.
    defer file.Close()

    // Step 5: Read the file contents into a byte slice.
	// Reading the file into a byte slice is a common practice because:
    // - It allows you to treat the file content as raw data, suitable for various data types.
    // - Provides flexibility to handle text, binary, or any kind of data without assumptions.
    // - Is more efficient for reading chunks of data from files, especially for large files.
    fileData, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Step 6: Convert the byte slice to a string and print it.
    fileContent := string(fileData)
    fmt.Println("File content:")
    fmt.Println(fileContent)
}