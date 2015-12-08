package main

import "fmt"

func main() {
	// tag::main[]
	var red uint8 = 255
	red++
	fmt.Println(red) //<1>

	var number int8 = 127
	number++
	fmt.Println(number) //<2>
	// end::main[]
}
