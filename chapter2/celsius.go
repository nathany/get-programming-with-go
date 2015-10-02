package main

import "fmt"

func main() {
	// tag::main[]
	type Celsius float64 //<1>

	var temperature Celsius = 20

	fmt.Println(temperature)
	// end::main[]
}
