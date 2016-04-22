package main

import "fmt"

func main() {
	var age = 39
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
	var command = "walk outside"
	var exit = (command == "walk outside")

	fmt.Println("Walk outside the cave?", exit)
}
