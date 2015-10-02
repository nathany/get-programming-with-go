package main

import "fmt"

func main() {
	// tag::main[]
	launch := false

	var countdown string

	if launch {
		countdown = "Launch in T minus 10 seconds."
	}
	fmt.Println(countdown)
	// end::main[]
}
