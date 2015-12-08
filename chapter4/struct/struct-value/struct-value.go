package main

import "fmt"

func main() {
	type point struct {
		lat  float64
		long float64
	}

	// tag::main[]
	bradbury := point{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.01 //<1>

	fmt.Println(bradbury, curiosity) //<2>
	// end::main[]
}
