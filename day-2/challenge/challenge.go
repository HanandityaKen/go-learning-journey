package main

import "fmt"

func main() {
	a := 10
	b := 3

  var sum int
  sum = a + b
  
  // TODO: Calculate and store the difference (a - b) in a variable named 'difference'
  var difference int
  difference = a - b
  
  var product int
  product = a * b
  
  var quotient int
  quotient = a / b
  
  var remainder int
  remainder = a % b
  
  fmt.Println("Sum:", sum)
  fmt.Println("Difference:", difference)
  fmt.Println("Product:", product)
  fmt.Println("Quotient:", quotient)
  fmt.Println("Remainder:", remainder)
}