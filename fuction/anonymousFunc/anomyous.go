package main

import "fmt"

func main() {
	add := func(x int, y int) int {
		return x + y
	}
	sum := add(2, 3)
	fmt.Println(sum)
}
