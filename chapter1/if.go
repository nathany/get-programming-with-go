package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(3) //<1>

	if num == 0 { //<2>
		fmt.Println("Space Adventures")
	} else if num == 1 { //<3>
		fmt.Println("SpaceX")
	} else { //<4>
		fmt.Println("Virgin Galactic")
	}
}
