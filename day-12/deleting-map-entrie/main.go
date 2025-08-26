package main

import "fmt"

func main() {
	inventory := map[string]int{
		"pen": 15,
		"pencil": 10,
		"paper": 25,
		"eraser": 5,
	}

	fmt.Println("Initial inventory:", inventory)

	delete(inventory, "pencil")

	fmt.Println("New inventory:", inventory)

	_, exists := inventory["pencil"]
	
	fmt.Println("Does 'pencil' exist in inventory?", exists)
}