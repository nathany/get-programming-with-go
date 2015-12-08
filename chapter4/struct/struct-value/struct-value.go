package main

import "fmt"

func main() {
	type point struct {
		lat  float64
		long float64
	}
	bradbury := point{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.01

	fmt.Println(bradbury, curiosity)
}
