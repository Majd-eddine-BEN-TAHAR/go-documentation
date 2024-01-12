package main

import (
	"fmt"
	"net/http"
	"validate/internal/handler"
	"validate/internal/validator"
)

func main() {
    validator.Setup()

    http.HandleFunc("/car", handler.CarHandler)
	http.HandleFunc("/user", handler.UserHandler)
    // Add more routes here

    fmt.Println("Server starting on port 8000...")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}