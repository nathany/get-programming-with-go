package main

import "fmt"

type report struct {
	sol       int
	high, low float64
	lat, long float64
}

func main() {

	report := report{sol: 15, high: -1.0, low: -78.0, lat: -4.5895, long: 137.4417}
	fmt.Printf("%+v\n", report)
}
