package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(3) //#A

	if num == 0 { //#B
		fmt.Println("Space Adventures")
	} else if num == 1 { //#C
		fmt.Println("SpaceX")
	} else { //#D
		fmt.Println("Virgin Galactic")
	}
}

// #A Generate a random number from 0-2
// #B If the random num is equal to 0
// #C Otherwise, if num is equal to 1
// #D Or, if nothing else
