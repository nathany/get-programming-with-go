package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// a random distance to Mars (km)
	var distance = rand.Intn(346000000) + 56000000
	fmt.Println(distance)
}
