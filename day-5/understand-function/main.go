package main

import "fmt"

func greet(name string) string {
	return "Hello " + name + "!"
}

func main() {
	name := "Gopher"
	message := greet(name)
	fmt.Println(message)
}