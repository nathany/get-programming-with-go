package main

import "fmt"

func main() {
	// tag::main[]
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
	// end::main[]

	// tag::blankid[]
	for _, c := range question {
		fmt.Printf("%c ", c) //<1>
	}
	// end::blankid[]
}
