package main

import ( //<1>
	"fmt"
	"math/rand" //<2>
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
