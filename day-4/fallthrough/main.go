package main

import "fmt"

func main() {
	// A variable representing a day of the week
	day := "Monday"
	
	// Switch statement to check the day
	switch day {
	case "Monday":
		fmt.Println("It's a weekday!")
		// TODO: Add the fallthrough keyword here to continue to the next case
        fallthrough
	case "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Time to work!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day!")
	}
}
