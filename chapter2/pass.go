package main

import (
	"fmt"
	"strings"
)

func main() {
	boardingPass := "InSight Mission to Mars"
	mars := strings.Contains(boardingPass, "Mars")

	if mars {
		fmt.Printf("Boarding pass: %v\n", boardingPass)
	}

	if strings.Contains(boardingPass, "Mars") {
		fmt.Printf("Mars: %v", boardingPass)
	} else {
		fmt.Printf("Not Mars: %v", boardingPass)
	}
}
