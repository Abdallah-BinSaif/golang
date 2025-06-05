package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var age int = 17
	if age < 18 {
		fmt.Println("You are a minor.")
	}
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5:", x)
	} else {
		fmt.Println("x is not greater than 5:", x)
	}
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Println("Successfully converted to int: ", num)
	} else {
		fmt.Println("Error converting string:", err)
	}
	dayOfWeek := "Wednesday"
	switch dayOfWeek {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	rand.Seed(time.Now().UnixNano())
	switch num := rand.Intn(10); {
	case num < 5:
		fmt.Printf("%d is less than 5.\n", num)
	case num == 5:
		fmt.Printf("%d is exactly 5\n", num)
	default:
		fmt.Printf("%d is greater than 5\n", num)
	}

}
