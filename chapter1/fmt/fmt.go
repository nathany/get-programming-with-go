package main

import "fmt"

func main() {
	// tag::simple[]
	var weight = 210.0 * 0.3783
	var age = 38 * 365 / 687

	fmt.Printf("My weight on the surface of Mars is %v lbs,", weight)
	fmt.Printf(" and I would be %v years old.\n", age)
	// end::simple[]

	// tag::multiple[]
	fmt.Printf("weight %v, age %v\n", weight, age)
	// end::multiple[]

	// tag::width[]
	fmt.Printf("$%4v\n", 94)  //<1>
	fmt.Printf("$%4v\n", 100) //<2>
	// end::width[]

}
