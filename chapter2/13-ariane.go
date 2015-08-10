package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32767

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// out of range
	}

	var h int16 = int16(bh) //#A
	fmt.Println(h)
}

// #A TODO: add rocket science
