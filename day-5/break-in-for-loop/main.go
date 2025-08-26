package main 

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Checking number: %d\n", i)

		if i == 3 {
			fmt.Printf("Found it: %d!\n", i)
			break 
		}
	} 

	fmt.Println("Search complete.")
}