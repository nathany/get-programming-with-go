package main

import "fmt"

func main() {
	// tag::main[]
	type point struct {
		lat  float64
		long float64
	}

	var spirit point //<1>
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity point //<1>
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity) //<2>
	// end::main[]
}
