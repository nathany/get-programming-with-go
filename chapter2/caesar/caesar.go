package main

import "fmt"

func main() {
	// tag::main[]
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c) //<1>
	// end::main[]

	c = 'x'
	c = c + 3
	// tag::wrap[]
	if c > 'z' {
		c = c - 26
	}
	// end::wrap[]
	fmt.Printf("%c", c)
}
