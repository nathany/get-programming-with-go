package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch speed := rand.Intn(15) + 16; speed {
	case 16:
		fmt.Printf("Slow going at %v km/s\n", speed)
	case 28, 29, 30:
		fmt.Printf("Quickly now at %v km/s\n", speed)
	default:
		fmt.Println(speed, "km/s")
	}
}
