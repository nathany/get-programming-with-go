package main

import "fmt"

func main() {
	// tag::main[]
	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune) //<1>
	// end::main[]

	// tag::neptune[]
	neptune = "🍨"
	fmt.Println(tune) //<1>
	// end::neptune[]

	// neptune[0] = 'H' // cannot assign to neptune[0]

	// tag::bytes[]
	question := "¿Cómo estás?"
	fmt.Println(question[:6]) //<1>
	// end::bytes[]

}
