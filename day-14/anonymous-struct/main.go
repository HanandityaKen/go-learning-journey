package main

import (
	"fmt"
)

func main() {
	book := struct {
		title		string
		pages		int
	} {
		title: "The Go Programming Language",
		pages: 380,
	}

	fmt.Printf("Book: %s, Pages: %d\n", book.title, book.pages)
}