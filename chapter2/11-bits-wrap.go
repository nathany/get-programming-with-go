package main

import "fmt"

func main() {
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue)
	blue++
	fmt.Printf("%08b\n", blue)
}

// #A 11111111
// #B 00000000
