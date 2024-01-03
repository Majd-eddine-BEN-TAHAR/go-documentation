package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Person struct: Defines a person with a name and age.
// If you want to put your struct in a separate file and use it, make sure to start the struct name with a capital letter to make it exported. Also, remember that the fields inside the struct should also begin with capital letters to make them accessible from anotehr file. If you miss that, you won't be able to modify or setup those fields from outside the file where you declared them
// Struct tags are used here for JSON encoding/decoding.
// The `json:"name"` tag tells the JSON encoder and decoder that the JSON key for the Name field is "name".
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Greet method (pointer receiver): Modifies the Person struct's name.
// Using a pointer receiver (*Person) allows this method to modify the actual Person instance, not just a copy.
// This is useful for updating or changing the original data.
func (p *Person) Greet() string {
    p.Name = "Mr. " + p.Name
    return "Hello, my name is " + p.Name
}

// SetAge (value receiver): Tries to change the Person's age.
// With a value receiver (Person), this method works on a copy of the struct, leaving the original unchanged.
// Useful when you want to work with the data but not modify the original struct.
// look to Person => it does not has * which means it's a value receiver, so a copy 
func (p Person) SetAge(newAge int) {
    p.Age = newAge
    // The original Person's age remains unaffected.
}

// CompareAge method: Compares the age of two Person structs.
// A value receiver is used here as we only need to read, not modify, the struct's data.
func (p Person) CompareAge(other Person) bool {
    return p.Age == other.Age
}

// Employee struct: Demonstrates composition over inheritance.
// Instead of inheriting from Person, Employee is composed of Person.
// Go use composition (having a struct as a field) over inheritance (extending a struct like in other languages);it'sa very fundamental concept in golang:composition over inheritance;
type Employee struct {
    Person
    Salary int
}

// declare a Speaker interface: it will be used with the Person struct to ensure that a Person implements a speak method
// Any struct with the Speak method is a Speaker.
type Speaker interface {
    Speak() string
}

// Speak method for thw Person struct: Implementation of the Speaker interface for Person.
func (p Person) Speak() string {
    return "Hello, I am " + p.Name + " and I am a speaker."
}

// UpdateAndPrintAge: Demonstrates passing a pointer to a struct and dereferencing.
// the function accepts a pointer to Person, allowing direct modification of the struct.
// Go automatically dereferences struct pointers when accessing their fields.
func UpdateAndPrintAge(p *Person, newAge int) {
    p.Age = newAge // // this was an implicit dereferencing done automatically by Go; the equivalent explicit  dereferencing  will be (*p).Age = newAge
    fmt.Println(p.Name, "is now", p.Age, "years old.") // Automatic dereferencing
}

// NewPerson function: Acts as a constructor for the Person struct.
// It includes basic validation to ensure the fields are not empty.
// returns a pointer to the created person
// this is the used way that we should use
func NewPerson(name string, age int) (*Person, error) {
    if name == "" || age <= 0 {
        return nil, errors.New("invalid input: name cannot be empty and age must be positive")
    }
    return &Person{Name: name, Age: age}, nil
}

func main() {
    // Initializing a Person struct with named fields.
    var john = Person{Name: "John", Age: 30}

    // Demonstration use of a method with a pointer receiver.
    fmt.Println(john.Greet())

    // Attempting to modify a struct using a value receiver method.
    // The original struct remains unchanged.
    john.SetAge(31)
    fmt.Println("John's age after SetAge:", john.Age) // Will still be 30, not 31.

    // Demonstrating passing a pointer to a function and dereferencing automatically in the function(go to the function to see that).
    UpdateAndPrintAge(&john, 31)

    // Comparing structs.
    jane := Person{"Jane", 30}
    if john.CompareAge(jane) {
        fmt.Println("John and Jane are the same age.")
    }else {
        fmt.Println("John and Jane are of different ages.")
    }

    // Using structs with interfaces.
    var speaker Speaker = Person{Name: "Alice"}
    fmt.Println(speaker.Speak())

    // Creating an Employee instance to demonstrate composition.
    bob := Employee{Person: Person{Name: "Bob", Age: 25}, Salary: 50000}
    fmt.Println("Employee Name:", bob.Name)
    fmt.Println("Employee Salary:", bob.Salary)

    // Struct Tags and JSON
    // Encoding a struct to JSON using tags.
    jsonData, _ := json.Marshal(john)
    fmt.Println("JSON representation of John:", string(jsonData))

    // Decoding JSON back to a struct.
    var decodedPerson Person
    json.Unmarshal(jsonData, &decodedPerson)
    fmt.Println("Decoded Person from JSON:", decodedPerson)

    // Using the NewPerson constructor to create a new Person instance.
    person, err := NewPerson("Mike", 28)
    if err != nil {
        fmt.Println("Error creating person:", err)
    } else {
        fmt.Println("Created Person:", *person)
    }

    // Attempting to create a Person with invalid data.
    invalidPerson, err := NewPerson("", -1)
    if err != nil {
        fmt.Println("Error creating person:", err)
    } else {
        fmt.Println("Created Person:", *invalidPerson)
    }
}