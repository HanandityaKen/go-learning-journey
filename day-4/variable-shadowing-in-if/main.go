package main 

import "fmt"

func main() {
	message := "Hello from outer scope"

	isLoggedIn := true

	if isLoggedIn {
		message := "Hello from inner scoper"

		fmt.Println(message)
	}

	fmt.Println(message)
}