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

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

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
	mars := world{radius: 3389.5}
	earth := world{radius: 6371}

	london := newPoint(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newPoint(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	fmt.Printf("London to Paris %.2f km\n", earth.distance(london, paris))

	edmonton := newPoint(coordinate{53, 32, 0, 'N'}, coordinate{113, 30, 0, 'W'})
	ottawa := newPoint(coordinate{45, 25, 0, 'N'}, coordinate{75, 41, 0, 'W'})
	fmt.Printf("Hometown to Capital %.2f km\n", earth.distance(edmonton, ottawa))

	mountSharp := newPoint(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newPoint(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	fmt.Printf("Mount Sharp to Olympus Mons %.2f km\n", mars.distance(mountSharp, olympusMons))
}
