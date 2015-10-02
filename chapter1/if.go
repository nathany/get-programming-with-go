package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(3)

	if num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
