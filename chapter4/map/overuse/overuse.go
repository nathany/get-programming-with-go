package main

import "fmt"

func main() {
	// weather data from Curiosity Rover
	// tag::main[]
	weather := map[string]float64{
		"sol":        16,
		"ls":         159.0,
		"min_temp":   -76.0,
		"max_temp":   0.0,
		"pressure":   7.4,
		"wind_speed": 2.0,
	}
	// end::main[]

	fmt.Println(weather)
}
