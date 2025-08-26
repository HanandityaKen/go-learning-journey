package main

import "fmt"

func openFile(filename string) error {
	if filename == "data.txt" {
		return nil // Ga eror atau benar
	} else {
		return fmt.Errorf("file %s not found", filename)
	}
}

func main() {
	filename := "config.txt"

	err := openFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return // return agar code diluar if tidak jalan
	}

	fmt.Println("File opened successfuly")
}