package main

import "fmt"

func main() {
	// tag::main[]
	s := []string{}
	lastCap := cap(s) //<1>

	for i := 0; i < 10000; i++ {
		s = append(s, "An element")
		if cap(s) != lastCap { //<2>
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
	// end::main[]
}
