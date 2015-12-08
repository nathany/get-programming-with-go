package main

import "fmt"

func main() {
	// tag::main[]
	lat := -4.5895   //<1>
	long := 137.4417 //<2>
	// end::main[]

	fmt.Println(lat, long)
}

func distance(lat1, long1, lat2, long2 float64) float64 {
	return 0
}
