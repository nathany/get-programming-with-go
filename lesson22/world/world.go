package main

import (
	"fmt"
	"math"
)

// location with a latitude, longitude.
type location struct {
	lat, long float64
}

// world with a volumetric mean radius in kilometers
type world struct {
	radius float64
}

var mars = world{radius: 3389.5}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
}
