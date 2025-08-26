package main

import "fmt"

type Temperature float64

func main() {
	var roomTemp Temperature = 23.5

	fmt.Printf("Room temperature: %v\n", roomTemp)
	fmt.Printf("Type of temperature: %T\n", roomTemp)
}