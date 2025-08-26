package main

import "fmt"

func main () {
	fruits := []string{"apple", "banana", "orange"}

	fruits = append(fruits, "grape", "kiwi")

	fmt.Println("My fruits collection:", fruits)
}