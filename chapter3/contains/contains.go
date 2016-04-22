package main

import (
	"fmt"
	"strings"
)

func main() {
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("Walk outside the cave?", exit)
}
