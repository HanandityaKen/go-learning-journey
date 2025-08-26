package main

import "fmt"

func main() {
	fruits := [5]string{"Apple", "Banana", "Orange", "Grape", "Mango"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

}