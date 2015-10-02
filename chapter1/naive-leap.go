package main

import "fmt"

func main() {
	// tag::main[]
	year := 2016
	if year%4 == 0 {
		fmt.Println("Look before you leap.")
	}
	// end::main[]
}
