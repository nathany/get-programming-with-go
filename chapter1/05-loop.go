package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0

	for count < 10 { //#A
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count = count + 1 //#B
	}
}

// #A Repeat this so long as count is less than 10
// #B Increment count, otherwise it will loop forever
