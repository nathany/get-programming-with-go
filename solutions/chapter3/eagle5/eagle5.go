package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch rand.Intn(7) {
	case 0, 1, 2:
		fmt.Println("Eagle 5")
	case 3:
		fmt.Println("Space Adventures")
	case 4:
		fmt.Println("SpaceX")
	case 5:
		fmt.Println("Virgin Galactic")
	case 6:
		fmt.Println("Blue Origin")
	default:
		fmt.Println("This was not foreseen")
	}
}
