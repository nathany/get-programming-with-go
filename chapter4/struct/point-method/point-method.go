package main

import (
	"fmt"
	"math"
)

// point with a latitude, longitude in decimal degrees.
type point struct {
	lat, long float64
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// String formats a point with latitude, longitude.
func (p point) String() string {
	lat := dms(p.lat, northSouth)
	long := dms(p.long, eastWest)
	return fmt.Sprintf("%v, %v", lat.String(), long.String())
}

// String formats a dms coordinate.
func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// dms converts decimal degrees to a coordinate in degrees, minutes, and seconds.
func dms(dd float64, hemi func(float64) rune) coordinate {
	h := hemi(dd)
	dd = math.Abs(dd)

	d := math.Floor(dd)
	m := math.Floor(60 * (dd - d))
	s := 3600 * (dd - d - m/60)
	return coordinate{d, m, s, h}
}

// northSouth returns the hemisphere for a latitude.
func northSouth(dd float64) rune {
	if dd < 0 {
		return 'S'
	}
	return 'N'
}

// eastWest returns the hemisphere for a longitude.
func eastWest(dd float64) rune {
	if dd < 0 {
		return 'W'
	}
	return 'E'
}

func main() {
	curiosity := point{-4.5895, 137.4417}
	fmt.Println(curiosity.String())
}
