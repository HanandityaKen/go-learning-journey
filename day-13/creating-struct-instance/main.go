package main

import "fmt"

type Person struct {
	name				string
	age 				int
	isStudent		bool
}

func main() {
	person := Person{name: "Alice", age: 25, isStudent: true}
	// Cara Gampang
	// person := Person{"Alice", 25, true}

	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)
	fmt.Printf("Is Student: %t\n", person.isStudent)
}