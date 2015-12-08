package main

import "fmt"

// point with a latitude, longitude.
type point struct {
	lat, long float64
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// tag::main[]
// newPoint from latitude, longitude d/m/s coordinates.
func newPoint(lat, long coordinate) point {
	return point{lat.decimal(), long.decimal()}
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	curiosity := newPoint(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}

// end::main[]
