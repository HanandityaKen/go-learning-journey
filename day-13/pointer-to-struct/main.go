package main

import "fmt"

type Person struct{
	name			string
	age				int
}

func main() {
	person := Person{
		"John",
		25,
	}
	
	personPtr := &person

	fmt.Println("Original person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	personPtr.name = "Jane"
	personPtr.age = 30

	fmt.Println("Updated person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	//Menampilkan alamat di struct -> pakai %p
	fmt.Printf("%p", personPtr)
}