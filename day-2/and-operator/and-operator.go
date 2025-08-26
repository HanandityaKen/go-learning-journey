package main

import "fmt"

func main() {
	hasExperience := true
	hasEducation := true

	isQualified := hasExperience && hasEducation

	fmt.Println("Does the candidat qualify for the job?", isQualified)
}