package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline Company Days Round-trip Price")
	fmt.Println("=======================================")

	var distance = 57600000
	var count = 0

	for count < 10 {
		var speed = rand.Intn(15) + 16          // 16-30 km/s
		var duration = distance / speed / 86400 // days
		var price = 20.0 + speed                // $ millions

		switch rand.Intn(3) {
		case 0:
			fmt.Print("Space Adventures  ")
		case 1:
			fmt.Print("SpaceX            ")
		case 2:
			fmt.Print("Virgin Galactic   ")
		}

		fmt.Printf("%4v ", duration)

		if rand.Intn(2) == 1 {
			fmt.Print("Round-trip ")
			price = price * 2
		} else {
			fmt.Print("One-way    ")
		}

		fmt.Printf("$%4v", price)
		fmt.Println()

		count = count + 1
	}
}
