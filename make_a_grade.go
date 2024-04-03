package main

import (
	"fmt"
)

func main() {
	var score float64

	// Prompt user for input
	fmt.Println("Enter the score:")
	fmt.Scanln(&score)

	// Calculate grade based on score
	grade := calculateGrade(score)

	// Output grade
	fmt.Println("Grade:", grade)
}

func calculateGrade(score float64) string {
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade
}
