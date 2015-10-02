package main

import "fmt"

func main() {
	// tag::main[]
	year := 2016
	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Leap for joy!")
	}
	// end::main[]
}
