package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if speed := rand.Intn(15) + 16; speed == 16 {
		fmt.Printf("Slow going %v km/s\n", speed)
	} else {
		fmt.Println(speed, "km/s")
	}
}
