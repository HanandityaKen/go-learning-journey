package main

import "fmt"

func main() {
	a := true
	b := false
	c := true

	result := a && (b || c)

	fmt.Println(result)
}