package main

import "fmt"

func main() {
	studentGrades := map[string]string{
		"John": "B+",
		"Emma": "A-",
		"Sarah": "A",
		"Mike": "C",
	}
	emmaGrade := studentGrades["Emma"]

	fmt.Println("Emma's grade is:", emmaGrade)
}