package main

import (
	"fmt"
	"errors"
)

func validateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username too short")
	}
	return nil
}

func main() {
	// Test with a valid username
	validName := "bob123"
	err1 := validateUsername(validName)
	if err1 != nil {
		fmt.Printf("%s is invalid: %v\n", validName, err1)
	} else {
		fmt.Printf("%s is valid!\n", validName)
	}
	
	// Test with an invalid username
	invalidName := "ab"
	err2 := validateUsername(invalidName)
	if err2 != nil {
		fmt.Printf("%s is invalid: %v\n", invalidName, err2)
	} else {
		fmt.Printf("%s is valid!\n", invalidName)
	}
}