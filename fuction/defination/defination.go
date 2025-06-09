package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}
func main() {
	a, b := 10, 11
	add(a, b)
}
