package main

import "fmt"

func main() {
  // Given variables
  score := 85
  grade := ""
	
	//coba coba
	// var score int
	// fmt.Print("Enter the score:")
	// fmt.Scan(&score)
  
  // TODO: Use an if statement to check if score is greater than or equal to 90
  // If true, set grade to "A"
  
  // TODO: Use an else if statement to check if score is greater than or equal to 80
  // If true, set grade to "B"
  
  // TODO: Use an else if statement to check if score is greater than or equal to 70
  // If true, set grade to "C"

  // TODO: Use an else statement to handle scores below 70
  // Set grade to "F"

  if score > 90 {
    grade = "A"
  } else if score >= 80 {
    grade = "B"
  } else if score >= 70 {
    grade = "C"
  } else {
    grade = "F"
  }
  
  // TODO: Print "Grade: B" using the grade variable
  fmt.Printf("Grade: %s\n", grade)
  
  // TODO: Use a switch statement with the grade variable to print a message
  // For "A", print "Excellent work!"
  // For "B", print "Good job!"
  // For "C", print "Satisfactory."
  // For any other grade, print "You need to improve."
  switch grade {
    case "A":
        fmt.Println("Excellent work!")
    case "B":
        fmt.Println("Good job!")
    case "C":
        fmt.Println("Satisfactory")
    default:
        fmt.Println("You need to improve.")
  }
}