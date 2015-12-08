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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants) //<1>
	// end::main[]

	// tag::gas[]
	fmt.Println(gasGiants[0]) //<1>
	// end::gas[]

	// tag::reslice[]
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice) //<1>
	// end::reslice[]

	// tag::ice[]
	iceGiantsMarkII := iceGiants //<1>
	iceGiants[1] = "🍨"
	fmt.Println(planets)                         //<2>
	fmt.Println(iceGiants, iceGiantsMarkII, ice) //<3>
	// end::ice[]
}
