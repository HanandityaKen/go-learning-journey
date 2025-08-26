package main

import "fmt"

func main() {
	// This represents the current hour in 24-hour format (0-23)
	timeOfDay := 14 // This is 2:00 PM
	
	// TODO: Complete the switch statement without an expression
	// to print the appropriate greeting based on the time of day
	switch {
	// Add your cases here
	case timeOfDay < 10:
        fmt.Println("Good morning!")
	case timeOfDay < 15:
        fmt.Println("Good afternoon!")
	case timeOfDay < 19:
        fmt.Println("Good evening!")
	case timeOfDay < 24:
        fmt.Println("Good night!")
	default:
		fmt.Println("Hello!")
	}
}