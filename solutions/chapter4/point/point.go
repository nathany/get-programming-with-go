package main

import "fmt"

func main() {
	type point struct {
		lat  float64
		long float64
	}

	var spirit point
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity point
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	var curiosity point
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(spirit, opportunity, curiosity)
}
