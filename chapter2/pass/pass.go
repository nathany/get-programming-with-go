package main

import (
	"fmt"
	"strings"
)

func main() {
	// tag::infer[]
	boardingPass := "InSight Mission to Mars"
	mars := strings.Contains(boardingPass, "Mars") //<1>

	if mars {
		fmt.Printf("Boarding pass: %v\n", boardingPass) //<2>
	}
	// end::infer[]

	// tag::direct[]
	if strings.Contains(boardingPass, "Mars") {
		fmt.Printf("Mars: %v", boardingPass)
	} else {
		fmt.Printf("Not Mars: %v", boardingPass)
	}
	// end::direct[]
}
