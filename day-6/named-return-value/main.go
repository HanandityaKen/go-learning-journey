package main

import "fmt"

func getPersonInfo() (name string, age int, isStudent bool) {
	name = "Alex"
	age = 25
	isStudent = true

	return
}

func main() {
	name, age, isStudent := getPersonInfo()
	fmt.Printf("Name : %s\nAge: %d\nStudent: %t\n", name, age, isStudent)
}