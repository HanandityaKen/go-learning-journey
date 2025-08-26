package main

import "fmt"

type Person struct{
	Name				string
	Age					int
	IsEmployed	bool
}

func main() {
	var alice Person
	
	alice = Person{"Alice", 28, true}

	fmt.Printf("Name: %s\n", alice.Name)
	fmt.Printf("Age: %d\n", alice.Age)
	fmt.Printf("Is Employed: %t\n", alice.IsEmployed)
}