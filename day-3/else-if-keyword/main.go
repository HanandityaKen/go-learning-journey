package main

import "fmt"

func main() {
	// This variable contains the current weather condition
	weather := "rainy"
	
	// TODO: Complete the if-else if-else statement to print the appropriate message
	if weather == "sunny" {
		fmt.Println("Wear sunscreen!")
	} else if weather == "rainy" {
		fmt.Println("Bring umberella!")
	} else if weather == "cold" {
		fmt.Println("Wear a jacket!")
	} else {
		fmt.Println("Check the forecast")
	}
	// Add your code here
}