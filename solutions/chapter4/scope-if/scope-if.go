package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if roundTrip := rand.Intn(2); roundTrip == 1 {
		speed := rand.Intn(15) + 16
		fmt.Printf("Round trip (%v) going %v km/s\n", roundTrip, speed)
	} else {
		// speed isn't defined here
		fmt.Println(roundTrip)
	}
	// roundTrip isn't defined here
}
