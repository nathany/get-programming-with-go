package main

import "fmt"

func main() {
	// tag::main[]
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895   //<1>
	curiosity.long = 137.4417 //<1>

	fmt.Println(curiosity.lat, curiosity.long) //<2>
	fmt.Println(curiosity)                     //<3>
	// end::main[]
}
