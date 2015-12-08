package main

import "fmt"

func main() {
	// tag::main[]
	var bh float64 = 32767
	var h int16 = int16(bh) //<1>
	fmt.Println(h)
	// end::main[]
}
