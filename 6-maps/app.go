package main

import "fmt"

func main()  {
/*  
    Maps vs. Structs in Go:
    - Maps are dynamic collections of key-value pairs. They are similar to JavaScript objects but with defined strict types from the begining(you cannot have string and int value types in the same map),    
    - Unlike JavaScript, the type of both the keys and the values must be defined and consistent within a map.
    - This strict typing means you cannot directly replicate JavaScript's flexibility in objects, like `const obj = {"name": "john", "age": 24}`.
    - Structs in Go are more like defined blueprints for objects. They have a fixed shape and can't have properties added or removed at runtime.

    
    Use Maps When:
        - You need a dynamic key-value store: Maps are perfect for scenarios where keys can change over time with adding and deletion of them.
        - You require quick lookup: Maps provide fast retrieval based on keys.
        - The keys and values are similar: All keys are of the same type, and all values are of the same type.
        - You need flexibility in the data structure: Maps are more flexible than structs as they can grow and shrink dynamically.
    
    Use Structs When:
        - You have a fixed set of fields: Use structs when you know in advance what the fields (and their types) will be, and these fields are unlikely to change frequently.
        - You need to represent an entity or object: For example, a Person struct with Name, Age, and Address fields.
        - You require methods: Structs allow you to define methods, making them suitable for object-oriented programming.
        - You need type safety: Structs ensure that each field is of a specified type.
    */


    /*  
        Creating a Map with a Map Literal:

        - You can initialize a map with a map literal, which allows you to define the map and its initial values simultaneously.
        - The syntax is: map[KeyType]ValueType{key1: value1, key2: value2, ...}.
        - This creates a fully functional map, to which you can add more key-value pairs, or modify or delete existing ones.
    */

    // Creating and initializing a map with city populations
    cityPopulation := map[string]int{
        "New York":    8419000,
        "Los Angeles": 3971000,
        "Chicago":     2716000,
    }

    // Printing the created map
    fmt.Println("City populations:", cityPopulation)

    // Modifying the map by adding a new city
    cityPopulation["Houston"] = 2325500
    fmt.Println("Updated city populations:", cityPopulation)

    /*  
    Creating a Map with the make Function:

    - The make function provides another way to initialize a map in Go. Unlike map literals, make is used to create an empty map.
    - The syntax for creating a map with make is: make(map[KeyType]ValueType, optionalCapacity).
    - The optionalCapacity parameter is not a limit; it's just a hint to the underlying implementation to anticipate the number of elements the map will have.
    - Using make is particularly useful when you know the approximate size of the map in advance. It can optimize memory allocation and potentially improve performance.
    - Unlike map literals, make does not allow you to initialize the map with predefined key-value pairs. It's ideal for situations where the size is known but the elements are not.
    - Example => m := make(map[string]int, 100) // Creates a map with an anticipated size of 100 entries.
*/

    userRoles := make(map[string]string, 5) // Creating a map with string keys and string values

    // Adding elements to the map:
    userRoles["alice"] = "admin" // Associating "admin" role with "alice"
    userRoles["bob"] = "user"   // Associating "user" role with "bob"

    // Accessing elements in the map:
    role := userRoles["alice"] // Accessing the role for "alice"
    fmt.Println("Alice's role:", role) // Printing Alice's role

    // Checking if a key exists in the map:
    _, exists := userRoles["carol"] // Checking if "carol" is in the map
    fmt.Println("Does Carol have a role?:", exists) // Printing whether Carol has a role

    // Iterating over a map in Go:
    for user, role := range userRoles {
        fmt.Printf("User: %s, Role: %s\n", user, role) // Printing each user and their role
    }

    // Deleting an element from the map:
    delete(userRoles, "bob") // Removing "bob" from the map
    fmt.Println("User roles after deleting Bob:", userRoles) // Printing the map after deletion

}