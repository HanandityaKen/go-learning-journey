package main

import "fmt"

func main() {
	rows := 3
	columns := 3

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}