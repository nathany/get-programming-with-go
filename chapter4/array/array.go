package main

import "fmt"

func main() {
	// tag::declare[]
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	// i := 8
	// planets[i] = "Pluto"
	// pluto := planets[i]
	// fmt.Println(pluto)
}
