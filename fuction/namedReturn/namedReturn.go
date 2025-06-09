package main

import "fmt"

func divide(numerator int, denominator int) (quotient int, remainder int) {
	if denominator == 0 {
		fmt.Println("Error: Division by Zero !")
		return 0, 0
	}
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func main() {
	q, r := divide(11, 3)
	fmt.Printf("11/3=Quotient: %d, Reminder:%d\n", q, r)
	divide(5, 0)
}
