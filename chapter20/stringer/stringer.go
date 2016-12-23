package main

import "fmt"

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long float64
}

// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}
