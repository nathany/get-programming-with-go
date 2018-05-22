package main

import "fmt"

func main() {
	// add a number larger than one
	var red uint8 = 255
	red += 2
	fmt.Println(red)

	var number int8 = 127
	number += 3
	fmt.Println(number)
	// wrap the other way
	red = 0
	red--
	fmt.Println(red)

	number = -128
	number--
	fmt.Println(number)
	// wrapping with a 16-bit unsigned integer
	var green uint16 = 65535
	green++
	fmt.Println(green)
}
