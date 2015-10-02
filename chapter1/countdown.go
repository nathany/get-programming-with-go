package main

import "fmt"

func main() {
	// tag::main[]
	var count = 10

	for count > 0 {
		fmt.Println(count)
		count = count - 1
	}
	// end::main[]
}
