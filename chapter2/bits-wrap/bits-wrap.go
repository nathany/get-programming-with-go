package main

import "fmt"

func main() {
	// tag::main[]
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue) //<1>
	blue++
	fmt.Printf("%08b\n", blue) //<2>
	// end::main[]
}
