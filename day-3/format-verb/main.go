package main

import "fmt"

func main() {
	// These variables are already defined for you
	age := 25
	height := 5.9
	isStudent := true
	name := "Alex"
	
	// TODO: Use fmt.Printf to print the following sentence with the correct format verbs:
	// "Name: Alex, Age: 25, Height: 5.9, Student: true"
	
	// Write your code below:
    fmt.Printf("Name: %v, Age: %d, Height: %f, Student: %t", name, age, height, isStudent)
}