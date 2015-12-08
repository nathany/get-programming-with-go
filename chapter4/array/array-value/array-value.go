package main

import "fmt"

func main() {
	// tag::main[]
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsMarkII := planets //<1>

	planets[2] = "💥" //<2>

	fmt.Println(planets)       //<3>
	fmt.Println(planetsMarkII) //<4>
	// end::main[]
}
