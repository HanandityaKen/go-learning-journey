package main

import "fmt"

func main() {
	favoriteColors := make(map[string]string)

	favoriteColors["Alice"] = "Blue"
	favoriteColors["Bob"] = "Green"
	favoriteColors["Charlie"] = "Red"

	fmt.Println("Favorite Colors:", favoriteColors)

	fmt.Println("Bob's favorite color is", favoriteColors["Bob"])
}