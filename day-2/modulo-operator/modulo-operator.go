package main

import "fmt"

func main() {
	number := 17
	divisor := 5

	var remainder int

	remainder = number % divisor

	fmt.Println("The remainder when", number, "is divided by", divisor, "is:", remainder)
}