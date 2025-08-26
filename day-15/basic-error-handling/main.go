package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return x / y, nil
}

func main() {
	numerator := 10.0
	denominator := 0.0

	result, err := divide(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}