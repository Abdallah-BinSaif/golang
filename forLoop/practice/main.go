package main

import (
	"fmt"
	"time"
)

/**

1. Standard for loop
2. for loop as a while loop
3. Infinite for loop
4. for .... range loop
5. break and continue statement

*/

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d \n", i)
	}

	// for loop as while loop
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Printf("Current sum : %d \n", sum)
	}

	// Infinit for loop
	counter := 0
	for {
		fmt.Printf("Infinite loop iteration %d \n", counter)
		counter++
		time.Sleep(500 * time.Millisecond)
		if counter == 3 {
			fmt.Println("Breaking out of infinite loop")
			break
		}
	}

	// for .... range loop
	numbers := []int{10, 20, 30, 40, 50, 60}
	for index, value := range numbers {
		fmt.Printf("Index: %d value : %d\n", index, value)
	}

	// break and continue
	fmt.Println("Example with continue: ")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Skipping iteration 2 with continue")
			continue
		}
		fmt.Printf("Iteration in continue no. : %d\n", i)
	}
	fmt.Println("Example with break : ")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Breaking at Iteration 2")
			break
		}
		fmt.Printf("Iteration in continue no. : %d\n", i)
	}
}
