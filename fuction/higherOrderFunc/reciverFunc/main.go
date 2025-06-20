package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (usr User) userDetails(a int) {
	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
	fmt.Println("A : ", a)
}

func main() {
	user1 := User{
		Name: "abdallah",
		Age:  34,
	}
	user2 := User{
		Name: "abdur rahman",
		Age:  41,
	}
	// userDetails(user1)
	// userDetails(user2)
	user1.userDetails(11)
	user2.userDetails(3)
}
