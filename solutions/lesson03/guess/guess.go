package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 42

	for {
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v is too small.\n", n)
		} else if n > number {
			fmt.Printf("%v is too big.\n", n)
		} else {
			fmt.Printf("You got it! %v\n", n)
			break
		}
	}
}
