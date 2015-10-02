package main

import "fmt"

func main() {
	// tag::main[]
	var green uint8 = 3
	fmt.Printf("%08b\n", green) //<1>
	green++
	fmt.Printf("%08b\n", green) //<2>
	// end::main[]
}
