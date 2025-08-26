package main

import "fmt"

func main() {
	var ptr *int

	if ptr == nil {
		fmt.Println("The pointer is nil")
	}

	if ptr != nil {
			fmt.Println("The pointer is not nill")
	}

	fmt.Println("Program completed")
}