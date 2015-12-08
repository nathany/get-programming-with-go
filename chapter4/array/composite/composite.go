package main

import "fmt"

func main() {
	// tag::main[]
	planets := [...]string{ //<1>
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune", //<2>
	}
	// end::main[]

	fmt.Println(planets)
	fmt.Println(len(planets))
}
