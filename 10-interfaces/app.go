package main

import "fmt"

// Speaker interface defines a set of rules for types.
// A type that implements the Speaker interface must have a Speak method.
// the naming convention when an interface has only one method is to take the name of the method inside it and add to it (er) and that would be the name of the interface
// In Go, explicit declaration of interface implementation is not required.
// A struct is considered to implement an interface if it defines all the methods the interface requires.
// Go uses "implicit interface implementation," meaning a struct satisfies an interface as long as it implements all its methods, without explicit declaration.
type Speaker interface {
    Speak() string
}

// Identifiable interface embeds the Speaker interface and adds an Identity method.
// This demonstrates interface embedding in Go, where one interface includes the methods of another.
// Types implementing Identifiable must implement both Speak and Identity methods.
type Identifiable interface {
    Speaker
    Identity() string
}

// Dog struct represents a dog with a Name attribute.
type Dog struct {
    Name string
}

// Cat struct represents a cat with a Name attribute.
type Cat struct {
    Name string
}

// Speak method for Dog, enabling it to fulfill the Speaker interface.
func (d Dog) Speak() string {
    return "My name is " + d.Name
}

// Identity method for Dog, fulfilling the Identifiable interface.
func (d Dog) Identity() string {
    return "Dog"
}

// Speak method for Cat, allowing it to fulfill the Speaker interface.
func (c Cat) Speak() string {
    return "My name is " + c.Name + ", and I say meow"
}

// Identity method for Cat, fulfilling the Identifiable interface.
func (c Cat) Identity() string {
    return "Cat"
}

// Introduce is a function that demonstrates the use of the Identifiable interface.
// It accepts any type that implements the Identifiable interface and invokes its Speak and Identity methods.
func Introduce(i Identifiable) {
    fmt.Printf("I am a %s: %s\n", i.Identity(), i.Speak())
}

// Describe is a function that demonstrates the use of the empty interface (interface{}), it's the any type in go.
// It accepts a parameter of any type and uses a type switch to handle different types.
// you can use (any) instead of interface{}, it's the same an it means the type can be (string,int,slice...)
func Describe(i interface{}) {  
    // Type switch to handle different types stored in interface{}.
    switch entityType := i.(type) {
    case Dog:
        fmt.Printf("Dog with Name: %s\n", entityType.Name)
    case Cat:
        fmt.Printf("Cat with Name: %s, says meow\n", entityType.Name)
    default:
        fmt.Printf("Unhandled type: %T\n", entityType)
    }
}

func main() {
    myDog := Dog{Name: "Buddy"}
    Introduce(myDog)  // Expected output: "I am a Dog: My name is Buddy"

    myCat := Cat{Name: "Whiskers"}
    Introduce(myCat)  // Expected output: "I am a Cat: My name is Whiskers, and I say meow"

    // Using Describe with empty interface (interface{})
    // This demonstrates how Describe can handle different types.
    Describe(myDog)
    Describe(myCat)
    Describe("Just a string")
}