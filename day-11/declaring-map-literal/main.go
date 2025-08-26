package main

import "fmt"

func main() {
	capitals := map[string]string{
		"France": "Paris",
		"Japan": "Tokyo",
		"Brazil": "Brasília",
	}

	fmt.Println(capitals)

	fmt.Println("The capital of Japan is:", capitals["Japan"])
}