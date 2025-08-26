package main

import "fmt"

func main() {
	homeworkComplete := false
	extraCreditComplete := true

	passesClass := homeworkComplete || extraCreditComplete

	fmt.Println("Homework complete:", homeworkComplete)
	fmt.Println("Extra credit complete:", extraCreditComplete)
	fmt.Println("Student passes class:", passesClass)
}