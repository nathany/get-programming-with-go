package main

import "fmt"

func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 154.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 40*365/687)
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 154.0)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
