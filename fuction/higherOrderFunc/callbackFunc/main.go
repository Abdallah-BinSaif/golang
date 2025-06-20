package main

import "fmt"

func paintWall(getColor func() string) {
	fmt.Println("Starting work for color the wall")
	fmt.Println("Prepare the room and Lay down the drop cloths")
	color := getColor()
	fmt.Println("")
	fmt.Println("Mixing the color")
	fmt.Printf("We did the room %s color", color)
}

func green() string {
	return "Green"
}
func blue() string {
	return "blue"
}
func main() {

	paintWall(green)
	fmt.Println("")
	fmt.Println("green room done")
	fmt.Println("")

	paintWall(blue)
	fmt.Println("")
	fmt.Println("blue room done")
}
