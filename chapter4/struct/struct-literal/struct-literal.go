package main

import "fmt"

func main() {
	type point struct {
		lat, long float64
	}

	opportunity := point{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	elysium := point{long: 154.7}
	fmt.Println(elysium)
	spirit := point{-14.5684, 175.472636}
	fmt.Println(spirit)
	curiosity := point{-4.5895, 137.4417}

	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}
