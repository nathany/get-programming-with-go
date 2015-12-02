package main

import "fmt"

func main() {
	type point struct {
		lat, long float64
	}

	spirit := point{-14.5684, 175.472636}
	fmt.Printf("%v\n", spirit)

	opportunity := point{lat: -1.9462, long: 354.4734}
	fmt.Printf("%+v\n", opportunity)

	var airy0 point
	fmt.Printf("%+v\n", airy0)

	elysium := point{long: 154.7}
	fmt.Println(elysium)
}
