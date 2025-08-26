package main

import "fmt"

func main() {
	favoriteFruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Strawberry"}

	var numberOfFruits int
	numberOfFruits = len(favoriteFruits)

	fmt.Printf("There are %d fruits in the array\n", numberOfFruits)

}