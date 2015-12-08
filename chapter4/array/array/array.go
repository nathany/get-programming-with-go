package main

import "fmt"

func main() {
	// tag::main[]
	// tag::declare[]
	var planets [8]string
	// end::declare[]

	planets[0] = "Mercury" //<1>
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] //<2>
	fmt.Println(earth)  //<3>
	// end::main[]

	// tag::length[]
	fmt.Println(len(planets))     //<1>
	fmt.Println(planets[3] == "") //<2>
	// end::length[]

	// i := 8
	// planets[i] = "Pluto"
	// pluto := planets[i]
	// fmt.Println(pluto)
}
