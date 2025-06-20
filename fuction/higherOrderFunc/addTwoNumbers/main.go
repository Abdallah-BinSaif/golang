package main

import "fmt"

func outerFunc(num1 int, num2 int, op func(a int, b int) int) int {
	result := op(num1, num2)
	return result
}

func add(x int, y int) int {
	sum := x + y
	fmt.Println("Sum: ", sum)
	return sum
}

func main() {
	finalresult := outerFunc(30, 20, add)
	fmt.Println(finalresult)
}
