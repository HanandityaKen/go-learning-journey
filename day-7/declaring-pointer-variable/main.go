package main

import "fmt"

func main() {
	originalValue := 42
	pointerToValue := &originalValue

	secondPointer := pointerToValue
	// secondPointer := &pointerToValue

	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)

	originalValue = 100
	fmt.Printf("\nAfter changing original value to 100:\n")
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
	
}