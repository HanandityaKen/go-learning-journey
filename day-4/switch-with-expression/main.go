package main

import "fmt"

func main() {
	// The day name is already defined
	dayName := "Wednesday"
	
	// Variable to store the description
	var description string
	
	// TODO: Write a switch statement that sets the description based on dayName
	// For Monday: "Start of the week"
	// For Wednesday: "Midweek"
	// For Friday: "Almost weekend"
	// For Saturday or Sunday: "Weekend"
	// For any other day: "Regular day"

    switch dayName {
    case "Monday":
        description = "Start og the week"
    case "Wednesday":
        description = "Midweek"
    case "Friday":
        description = "Almost weekend"
    case "Sunday":
        description = "Weekend"
    default:
        description = "Regular day"
    }
	// Print the result
	fmt.Printf("%s is %s\n", dayName, description)
}