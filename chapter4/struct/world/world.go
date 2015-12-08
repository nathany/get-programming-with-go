package main

import (
	"fmt"
	"math"
)

// point with a latitude, longitude.
type point struct {
	lat, long float64
}

// world with a volumetric mean radius in kilometers
type world struct {
	radius float64
}

var mars = world{radius: 3389.5}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 point) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// destination calculates a new point.
func (w world) destination(p1 point, distance, direction float64) point {
	slat1, clat1 := math.Sincos(rad(p1.lat))
	sb, cb := math.Sincos(rad(direction))
	sd, cd := math.Sincos(distance / w.radius)

	lat2 := math.Asin(slat1*cd + clat1*sd*cb)
	long2 := rad(p1.long) + math.Atan2(sb*sd*clat1, cd-slat1*math.Sin(lat2))
	return point{deg(lat2), deg(long2)}
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// deg converts radians to degrees.
func deg(rad float64) float64 {
	return rad * 180 / math.Pi
}

func main() {
	spirit := point{-14.5684, 175.472636}
	opportunity := point{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
	bradbury := point{-4.5895, 137.4417}
	fmt.Println(bradbury)

	yellowknife := mars.destination(bradbury, .624, 90)
	fmt.Println(yellowknife)
}
