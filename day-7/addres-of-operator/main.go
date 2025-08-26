package main

import "fmt"

func main() {
	name := "Gopher"

	var namePointer *string

	namePointer = &name

	fmt.Printf("The value at that memory address is: %v\n", *namePointer) // Hasil "The value at that memory address is: Gopher"
	// fmt.Printf("The value at that memory address is: %v\n", namePointer) // Hasil "The value at that memory address is: 0xc0000300a0"
}