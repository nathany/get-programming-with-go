package main

import "fmt"

func main() {
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

	// tag::indices[]
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	// end::indices[]

	fmt.Println(terrestrial, gasGiants, iceGiants) //<1>

	// tag::all[]
	allPlanets := planets[:]
	// end::all[]

	fmt.Println(allPlanets) //<1>
}
