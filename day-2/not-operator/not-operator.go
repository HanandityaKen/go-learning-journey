package main

import "fmt"

func main() {
	isRaining := true
	result := !isRaining

	fmt.Println("Original value:", isRaining)
	fmt.Println("After using NOT operator:", result)
}