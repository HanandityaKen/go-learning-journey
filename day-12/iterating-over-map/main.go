for package main

import (
	"fmt"
	"sort"
)

func main() {
	studentGrades := map[string]string{
		"Alice": "A",
		"Bob": "B",
		"Charlie": "C",
		"David": "A-",
		"Eva": "C",
	}

	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("Student: %s, Grade: %s\n", name, studentGrades[name])
	}
}