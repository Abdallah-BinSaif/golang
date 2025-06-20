package main

import "fmt"

func main() {
	var numbers [5]int
	var names [3]string = [3]string{"abdallah", "abdurrahman", "anabia"}
	var fruites [10]string
	prices := [3]float32{10.0, 2.3, 8.3}
	cities := [...]string{"dhaka", "comilla"}

	fmt.Println(numbers)
	fmt.Println(names)
	fmt.Println(prices)
	fmt.Println(len(cities))
	fmt.Println(fruites)
}
