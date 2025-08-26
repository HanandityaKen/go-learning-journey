package main

import "fmt"

type Person struct {
	name			string
	age				int
	isStudent	bool
}

func main() {
	// john := Person{
	// 	name: "John Doe",
	// 	age: 25,
	// 	isStudent: true,
	// }

	john := Person{"John Doe", 25, true}

	fmt.Printf("Name: %s\n", john.name)
	
	fmt.Printf("Age: %d\n", john.age)

	fmt.Printf("Is Student: %t\n", john.isStudent)
}


