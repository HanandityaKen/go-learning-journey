package main

import (
	"errors"
	"fmt"
)

func devide(a, b int) (int, error) {
	if b == 0 {
		// return 0 kalo error
		return 0, errors.New("division zero")
	}
	// error nil kalau benar
	return a / b, nil
}

// Nyoba
func perkalian(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("tidak boleh kali 0")
	}

	return a * b, nil
}

func main() {
	//first case
	a, b := 10, 2

	result, err := devide(a, b)
	if err != nil {
		// kalo error nya engga nil, artinya kode salah
		fmt.Printf("Error: %s\n", err)
	} else {
		// kalo error nya nil, code benar
		fmt.Printf("Result: %d\n", result)
	}

	// second case
	a, b = 8, 0
	
	result, err = devide(a, b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Nyoba perkalian
	a, b = 2, 0

	result, err = perkalian(a, b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}