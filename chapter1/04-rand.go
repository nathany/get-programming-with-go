package main

import ( //#A
	"fmt"
	"math/rand" //#B
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}

// #A Import multiple packages with parenthesis
// #B math/rand is the import path for the rand package
