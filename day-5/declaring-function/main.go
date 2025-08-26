package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	name1 := "Alice"
	name2 := "Bob"

	message1 := greet(name1)
	message2 := greet(name2)

	fmt.Println(message1)
	fmt.Println(message2)
}