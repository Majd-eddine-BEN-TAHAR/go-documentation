package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func collectGoFilesDocs(rootFolder, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// packageName := "event_booking_api" // Set the package name

	err = filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".sum") && !strings.HasSuffix(path, ".mod") {
			relPath, err := filepath.Rel(rootFolder, path) // Get the relative path
			if err != nil {
				return err
			}

			fmt.Fprintf(writer, "\n%s\nFile: /%s\n%s\n", strings.Repeat("-", 20), relPath, strings.Repeat("-", 20))

			fileContent, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			writer.Write(fileContent)
			fmt.Fprint(writer, "\n"+strings.Repeat("-", 40)+"\n")
		}

		return nil
	})

	return err
}

func main() {
	rootFolder := "/home/yonkersleroy/Documents/go/go-course/25-event_booking_rest_api" // Root folder path
	outputFile := "collected_docs.txt"                                         // Output file name

	if err := collectGoFilesDocs(rootFolder, outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}