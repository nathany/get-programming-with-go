package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Departure  Spaceline Company Days Round-trip Price")
	fmt.Println("==================================================")

	var maxPrice = 80

	var distance = 0
	var departure = ""
	var count = 0

	for count < 10 {
		switch rand.Intn(3) {
		case 0:
			departure = "2016-05-22"
			distance = 75300000
		case 1:
			departure = "2018-07-27"
			distance = 57600000
		case 2:
			departure = "2020-10-13"
			distance = 62100000
		}

		var speed = rand.Intn(15) + 16          // 16-30 km/s
		var duration = distance / speed / 86400 // days
		var price = 20.0 + speed                // $ millions

		var roundTrip = rand.Intn(2)
		if roundTrip == 1 {
			price = price * 2
		}

		if price >= maxPrice {
			continue
		}

		fmt.Print(departure, " ")

		switch rand.Intn(3) {
		case 0:
			fmt.Print("Space Adventures  ")
		case 1:
			fmt.Print("SpaceX            ")
		case 2:
			fmt.Print("Virgin Galactic   ")
		}

		fmt.Printf("%4v ", duration)

		if roundTrip == 1 {
			fmt.Print("Round-trip ")
		} else {
			fmt.Print("One-way    ")
		}

		fmt.Printf("$%4v", price)
		fmt.Println()

		count = count + 1
	}
}
