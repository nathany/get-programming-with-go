package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// tag::main[]
	switch rand.Intn(3) {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("This was not foreseen")
	}
	// end::main[]
}
