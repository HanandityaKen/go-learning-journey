package main

import (
	"errors"
	"fmt"
)

func validateUsername(username string) (string, error) {
	if len(username) < 5 {
		return "", errors.New("username must be at least 5 character long")
	}

	return "Username is valid", nil
}

func main() {
	// Test with valid userame
	result, err := validateUsername("gopher123")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Test with invalid username
	result, err = validateUsername("go")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}