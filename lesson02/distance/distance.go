package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// a random distance to Mars (km)
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)
}
