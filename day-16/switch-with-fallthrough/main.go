package main

import "fmt"

func main() {
	var alertLevel string
	fmt.Println("Input alert level:")
	fmt.Scanln(&alertLevel)

	switch alertLevel {
		case "critical":
			fmt.Println("CRITICAL: System shutdown imminent!")
			fallthrough
		case "high":
			fmt.Println("HIGH: Immediate attention required!")
			fallthrough
		case "medium":
			fmt.Println("MEDIUM: Please review when possible")
			fallthrough
		case "low":
			fmt.Println("Low: Informational only")
	}
}