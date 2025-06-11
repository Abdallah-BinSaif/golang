package main

import "fmt"

var userName string = "abdallah"
var userPassword string = "password"

func greetings() {
	fmt.Println("--------- welcome ----------")
	fmt.Println("----------------------------")
}

func takeUserName() string {
	fmt.Println("UserName:- ")
	var inputUserName string
	fmt.Scanln(&inputUserName)
	return inputUserName
}

func takeUserPass() string {
	fmt.Println("UserPassword:- ")
	var inputUserPassword string
	fmt.Scanln(&inputUserPassword)
	return inputUserPassword
}
func main() {
	greetings()
	inputUserName := takeUserName()
	inputUserPassword := takeUserPass()
	if inputUserName == userName && inputUserPassword == userPassword {
		fmt.Println("Your user name and password match.")
		fmt.Println("You are authenticated")
	} else {
		fmt.Println("Your user name or password dont match")
		fmt.Println("You are not authenticated")
	}
}
