package main

import "fmt"

func main() {
    /*  
        Slices in Go:
            - Slices are dynamic arrays in Go. Unlike arrays, slices can grow or shrink in size.
            - Slices are more common in Go than arrays due to their flexibility.
            - Importantly, a slice is a reference to an underlying array(it's not an array but a refernce to an underlying array). Any modifications to the slice will also change the array it's referencing. And if you make a change in the original array, the slice also will change because it's a pointer to the array.
    */

    // Example of a Slice:
    var fruits []string = []string{"Apple", "Banana", "Cherry"}
    fmt.Println("Original Slice:", fruits)

    // Adding an element to a slice:
    fruits = append(fruits, "Date")
    fmt.Println("The Slice after adding an element:", fruits)

    // Slicing a slice:
    // This creates a slice including elements at index 1 and 2.
    // Note: 'subFruits' is a window to the 'fruits' slice(it means it's not another copy instead it's just a pointer with a start and end indexes to know your size). Any changes to 'subFruits' will be reflected in 'fruits' and vice versa, as they share the same underlying array becasue they share the same pointer
    subFruits := fruits[1:3] 
    fmt.Println("Sliced slice:", subFruits)

    // Modifying the sliced slice:
    // Changing the first element of 'subFruits' from 'Banana' to 'Blackberry' will affect the original slice
    subFruits[0] = "Blackberry"
    // Now, the change will reflect in the original 'fruits' slice as well.
    fmt.Println("Modified 'subFruits' slice:", subFruits)
    fmt.Println("Original 'fruits' slice after modification:", fruits)

    // Iterating over a slice:
    for i, fruit := range fruits {
        fmt.Printf("Element %d: %s\n", i, fruit)
    }

    // Copying a slice(a copy, so no modidfications to the original now)
    // Initialize a new slice 'copiedFruits' with the same length as 'fruits' using the 'make' function.
    copiedFruits := make([]string, len(fruits))
    // Use the 'copy' function to duplicate the elements from 'fruits' into 'copiedFruits'.
    // This function ensures that 'copiedFruits' has its own separate underlying array.
    // Changes to 'copiedFruits' will not affect 'fruits' and vice versa.
    copy(copiedFruits, fruits)
    fmt.Println("Copied slice:", copiedFruits)

    // Modifying the copied slice:
    copiedFruits[0] = "Avocado"
    // Observe that the original 'fruits' slice remains unchanged.
    fmt.Println("Modified 'copiedFruits' slice:", copiedFruits)
    fmt.Println("Original 'fruits' slice after modifying 'copiedFruits':", fruits)


    // Understanding the 'cap' function:
    // - The 'cap' function in Go returns the capacity of the slice, 
    //   which is the maximum number of elements it can hold without reallocating memory.
    // - When you append elements to a slice and its length exceeds its capacity,
    //   Go automatically allocates a new larger array and copies the elements to it.

    // Example of using the 'cap' function:
    fmt.Println("Capacity of the original 'fruits' slice:", cap(fruits))
    
    // Appending more elements to see how capacity changes:
    fruits = append(fruits, "Elderberry", "Fig", "Grape")
    fmt.Println("The Slice after adding more elements:", fruits)
    fmt.Println("Capacity of the 'fruits' slice after appending:", cap(fruits))

    // Observing capacity changes with a new slice:
    moreFruits := []string{"Honeydew", "Ivy Gourd"}
    fmt.Println("Initial 'moreFruits' slice:", moreFruits)
    fmt.Println("Initial capacity of 'moreFruits' slice:", cap(moreFruits))
    moreFruits = append(moreFruits, "Jackfruit")
    fmt.Println("Modified 'moreFruits' slice:", moreFruits)
    fmt.Println("Capacity of 'moreFruits' slice after appending:", cap(moreFruits))

    // Demonstrating the difference between length and capacity:
    fmt.Printf("Length vs Capacity of 'fruits': %d vs %d\n", len(fruits), cap(fruits))
    fmt.Printf("Length vs Capacity of 'moreFruits': %d vs %d\n", len(moreFruits), cap(moreFruits))


    // Understanding Slice Unpacking with 'append':
    // - To append all elements of one slice to another, Go requires the use of the '...' operator.
    // - This operator unpacks the slice, allowing each element to be appended individually.
    // - Without the '...', Go will not accept appending a complete slice directly as it treats the slice as a single entity.

    // Example of Slice Unpacking:
    // Initializing two slices
    vegetables := []string{"Carrot", "Potato", "Onion"}
    moreVegetables := []string{"Cucumber", "Lettuce", "Tomato"}

    // Printing the original slices
    fmt.Println("Original 'vegetables' slice:", vegetables)
    fmt.Println("Original 'moreVegetables' slice:", moreVegetables)

    // Attempting to append 'moreVegetables' to 'vegetables' using the '...' operator
    vegetables = append(vegetables, moreVegetables...)
    fmt.Println("Combined slice after unpacking 'moreVegetables' into 'vegetables':", vegetables)

    // Uncomment the line below to see the compiler error without the '...'
    // vegetables = append(vegetables, moreVegetables)
}