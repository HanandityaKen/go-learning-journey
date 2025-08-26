package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}

	for i, value := range fruits {
		fmt.Printf("%d. %s\n", i+1, value)
	}
}