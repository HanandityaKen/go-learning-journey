package main

import "fmt"

func main() {
	num1 := 10
	num2 := 10
	name1 := "Go"
	name2 := "Phyton"

	numbersEqual := num1 == num2

	namesNotEqual := name1 != name2

	fmt.Println("Are the numbers equal?", numbersEqual)
	fmt.Println("Are the names different?", namesNotEqual)
}