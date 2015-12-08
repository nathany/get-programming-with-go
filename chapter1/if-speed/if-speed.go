package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var speed = rand.Intn(15) + 16

	if speed == 16 {
		fmt.Println("Slow going 16 km/s")
	} else {
		fmt.Println(speed, "km/s")
	}
}
