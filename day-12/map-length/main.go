package main

import "fmt"

func main() {
	countries := map[string]string{
		"France": "Paris",
		"Japan": "Tokyo",
		"Brazil": "Brasilia",
		"Canada": "Ottawa",
		"Egypt": "Cairo",
	}

	countryCount := len(countries)

	fmt.Printf("The map contains %d countries\n", countryCount)
}