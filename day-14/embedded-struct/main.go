package main

import "fmt"

type Addres struct {
	Street		string
	City			string
	ZipCode		string
}

type Person struct {
	Name		string
	Age			int
	Addres
}

func main() {
	// person := Person {
	// 	Name: "Alice",
	// 	Age: 30,
	// 	Addres: Addres{
	// 		Street: "123 Main St",
	// 		City: "Wonderland",
	// 		ZipCode: "12345",
	// 	},
	// }

	person := Person{
		"Alice",
		30,
		Addres{
			"123 Main St",
			"Wonderland",
			"12345",
		},
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age", person.Age)

	fmt.Println("Street:", person.Street)
	fmt.Println("City:", person.City)
	fmt.Println("ZipCode:", person.ZipCode)
}