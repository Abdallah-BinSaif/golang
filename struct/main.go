package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{
		Name: "abdallah",
	}
	user2 := User{
		"abdurrahman", 25,
	}
	fmt.Println(user1.Name)
	fmt.Println(user2.Age)
}
