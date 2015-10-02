package main

import "fmt"

func main() {
	// tag::main[]
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "ºF\n") //<1>
	fmt.Print((9.0/5.0*celsius)+32, "ºF\n") //<1>
	// end::main[]
}
