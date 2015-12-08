package main

import (
	"fmt"
	"strconv"
)

func main() {
	// tag::main[]
	score := 128533

	str := "Score " + strconv.Itoa(score) //<1>
	fmt.Println(str)                      //<2>
	// end::main[]
}
