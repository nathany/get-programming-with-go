package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// point with a latitude, longitude in decimal degrees.
type point struct {
	Lat      float64 `json:"latitude"`
	Long     float64 `json:"longitude"`
	LatText  string  `json:"lat_text"`
	LongText string  `json:"long_text"`
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// String formats a point with latitude, longitude.
func (p point) String() string {
	lat := dms(p.Lat, northSouth)
	long := dms(p.Long, eastWest)
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
	curiosity := point{Lat: -4.5895, Long: 137.4417}
	curiosity.LatText = dms(curiosity.Lat, northSouth).String()
	curiosity.LongText = dms(curiosity.Long, eastWest).String()

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
