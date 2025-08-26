package main

import "fmt"

func main() {
	username := "bob"

	minLength := 5

	if len(username) < minLength {
		err := fmt.Errorf("invalid username: %s is too short, minimum length is %d", username, minLength)

		fmt.Println(err)
	} else {
		fmt.Println("Username is valid!")
	}

}