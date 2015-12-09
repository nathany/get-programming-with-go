package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Pi = 3.14159265358979323846264338327950288419716939937510582097494459

	pi64 := math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)
}
