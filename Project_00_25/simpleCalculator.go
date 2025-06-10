package main

import "fmt"

func addition(num1 int, num2 int) int {
	return num1 + num2
}

func sutraction(num1 int, num2 int) int {
	return num1 - num2
}

func multiplication(num1 int, num2 int) int {
	return num1 * num2
}

func division(num1 int, num2 int) int {
	return num1 / num2
}

func welcomeGreeting() {
	fmt.Println("_______Simple Calculator_______")
	fmt.Println("Here you can Addition Subtraction multiplication and Division")
}

func takeTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter 1st number")
	fmt.Scanln(&num1)

	fmt.Println("Enter 1st number")
	fmt.Scanln(&num2)

	return num1, num2
}

func takeParameter() string {
	var parameter string
	fmt.Println("1. Addition------> ( + )")
	fmt.Println("2. Subtraction------> ( - )")
	fmt.Println("3. Multiplication------> ( * )")
	fmt.Println("4. Divition------> ( / )")
	fmt.Println("Enter Perameter")
	fmt.Scanln(&parameter)
	return parameter
}

func main() {
	welcomeGreeting()
	num1, num2 := takeTwoNumbers()
	parameter := takeParameter()
	switch parameter {
	case "+":
		result := addition(num1, num2)
		fmt.Println(result)
	case "-":
		result := sutraction(num1, num2)
		fmt.Println(result)
	case "*":
		result := multiplication(num1, num2)
		fmt.Println(result)
	case "/":
		result := division(num1, num2)
		fmt.Println(result)
	default:
		fmt.Println("Invalid Parameter")
	}
}
