package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)

	fmt.Println("Length:", len(numbers))

	fmt.Println("Capacity:", cap(numbers))
}