package main

import "fmt"

func main() {
	var score int = 55
	var grade string
	if score >= 90 && score <= 100 {
		grade = "A"
	} else if score >= 80 && score < 90 {
		grade = "B"
	} else if score >= 70 && score < 80 {
		grade = "C"
	} else if score >= 60 && score < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	switch grade {
	case "A":
		fmt.Println("Score:", score, "Grade: ", grade)
	case "B":
		fmt.Println("Score:", score, "Grade: ", grade)
	case "C":
		fmt.Println("Score:", score, "Grade: ", grade)
	case "D":
		fmt.Println("Score:", score, "Grade: ", grade)
	case "F":
		fmt.Println("Score:", score, "Grade: ", grade)
	}
}
