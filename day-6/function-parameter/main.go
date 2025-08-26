package main

import "fmt"

func createGreeting(text1, text2 string) string {
	return text1 + ", " +text2 + "!"
	// return fmt.Sprintf("%s, %s", text1, text2)
}

func main() {
	message := createGreeting("Hello", "Gopher")
	fmt.Println(message)

	message1 := createGreeting("Welcome", "Friend")
	fmt.Println(message1)
}