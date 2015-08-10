package main

import "fmt"

func main() {
	var green uint8 = 3
	fmt.Printf("%08b\n", green) //#A
	green++
	fmt.Printf("%08b\n", green)
}

// #A 00000011
// #B 00000100
