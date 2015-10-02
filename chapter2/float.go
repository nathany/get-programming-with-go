package main

import "fmt"

func main() {
	// tag::main[]
	third := 1.0 / 3.0
	fmt.Println(third + third + third) //<1>

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank) //<2>
	// end::main[]
}
