package main

import (
	"errors"
	"fmt"
)

func GenerateError() error {
	return errors.New("something went wrong")
}

func main() {
	err := GenerateError()

	// tidak return 0 di func main
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
	}

	fmt.Println("Program completed successfully")
}