package main

import "fmt"

func multiplier(factor int) func(number int) int {
	return func(number int) int {
		return number * factor
	}
}

func main() {
	multiplyByTwo := multiplier(2)
	multiplyByThree := multiplier(3)

	fmt.Println(multiplyByTwo(7))
	fmt.Println(multiplyByThree(7))
}
