package main

import "fmt"

func main() {
	// tag::opportunity[]
	type point struct {
		lat, long float64
	}

	opportunity := point{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity) //<1>

	elysium := point{long: 154.7}
	fmt.Println(elysium) //<2>
	// end::opportunity[]

	// tag::spirit[]
	spirit := point{-14.5684, 175.472636}
	fmt.Println(spirit) //<1>
	// end::spirit[]

	// tag::fmtv[]
	curiosity := point{-4.5895, 137.4417}

	fmt.Printf("%v\n", curiosity)  //<1>
	fmt.Printf("%+v\n", curiosity) //<2>
	// end::fmtv[]
}
