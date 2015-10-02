package main

import "fmt"

func main() {
	// tag::main[]
	celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "ºF") //<1>
	// end::main[]
}
