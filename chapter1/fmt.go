package main

import "fmt"

func main() {
	var weight = 210.0 * 0.3783
	var age = 38 * 365 / 687

	fmt.Printf("My weight on the surface of Mars is %v lbs,", weight)
	fmt.Printf(" and I would be %v years old.\n", age)

	fmt.Printf("weight %v, age %v\n", weight, age)

	fmt.Printf("$%4v\n", 94)
	fmt.Printf("$%4v\n", 100)
}
