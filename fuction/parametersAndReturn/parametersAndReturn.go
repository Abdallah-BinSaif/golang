package main

import (
	"fmt"
	"strconv"
)

func getNumbers(num1, num2 int) (int, int) {
	sum := num1 + num2
	mult := num1 * num2
	return sum, mult
}
func parseInt(str1 string) (int, error) {
	num, err := strconv.Atoi(str1)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
func main() {
	a, b := 2, 3
	sum, mult := getNumbers(a, b)
	fmt.Println(sum)
	fmt.Println(mult)

	val, err := parseInt("123s")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Converted value", val)
	}
}
