package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red) //#A

	var number int8 = 127
	number++
	fmt.Println(number) //#B
}

// #A 0
// #B -128
