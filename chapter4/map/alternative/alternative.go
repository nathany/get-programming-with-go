package main

import "fmt"

func main() {
	// Weather data from Curiosity
	type Weather struct {
		sol              int
		ls               float64
		minTemp, maxTemp float64
		pressure         float64
		windSpeed        float64
	}

	weather := Weather{
		sol:       16,
		ls:        159.0,
		minTemp:   -76.0,
		maxTemp:   0.0,
		pressure:  7.4,
		windSpeed: 2.0,
	}

	fmt.Printf("%+v", weather)
}
