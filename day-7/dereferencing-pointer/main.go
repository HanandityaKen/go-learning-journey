package main

import "fmt"

func main() {
	message := "Hello, pointers!"

	messagePtr := &message
	
	value := *messagePtr

	fmt.Println("The pointers to the value:", value)
}