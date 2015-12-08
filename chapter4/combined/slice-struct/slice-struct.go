package main

import "fmt"

func main() {
	// tag::bad[]
	// tag::highs[]
	highs := []float64{0, -18, 0, 0, 0, -2}
	// end::highs[]
	lows := []float64{-76, -78, -73, -70, -73, -73}
	// end::bad[]

	// tag::good[]
	temperatures := []struct {
		high, low float64
	}{
		{high: 0, low: -76}, {high: -18, low: -78}, {high: 0, low: -73},
		{high: 0, low: -70}, {high: 0, low: -73}, {high: -2, low: -73},
	}

	fmt.Printf("high: %v, low: %v", temperatures[1].high, temperatures[1].low) //<1>
	// end::good[]

	_, _, _ = temperatures, highs, lows
}
