package main

import "fmt"

func getPersonInfo() (string, int, bool) {
	name := "Alex"
	age := 25
	isStudent := true

	return name, age, isStudent
}

func main() {
	name, age, isStudent := getPersonInfo()

	// Print the values
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
}