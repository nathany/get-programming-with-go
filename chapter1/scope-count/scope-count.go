package main

import "fmt"

func main() {
	// tag::main[]
	for count := 10; count > 0; count-- { //<1>
		fmt.Println("T minus", count)
	} //<2>
	fmt.Println("Liftoff!")

	for count := 1; count <= 10; count++ { //<3>
		fmt.Println("T plus", count)
	} //<4>
	// end::main[]
}
