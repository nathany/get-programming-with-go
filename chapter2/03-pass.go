package main

import (
	"fmt"
	"strings"
)

func main() {
	mark := "B+"
	pass := strings.ContainsAny(mark, "ABC") //#A

	fmt.Printf("Is %v a pass? %v\n", mark, pass) //#B

	if strings.ContainsAny(mark, "ABC") {
		fmt.Printf("%v is a pass", mark)
	} else {
		fmt.Printf("%v is a fail", mark)
	}
}

// #A Go infers the type based on the function being called
// #B Is B+ a pass? true
