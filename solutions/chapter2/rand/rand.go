package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// a random speed from 16-30 km/s
	var speed = rand.Intn(15) + 16
	fmt.Println(speed)
}
