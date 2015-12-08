package main

import "fmt"

func main() {
	// tag::main[]
	var happy rune = 128515
	var sad rune = 128546
	var bang byte = 33

	fmt.Println(happy, sad, bang) //<1>
	// end::main[]

	// tag::string[]
	fmt.Println(string(happy), string(sad), string(bang)) //<1>
	// end::string[]

	// tag::verb[]
	fmt.Printf("%c %c %c\n", happy, sad, bang) //<1>
	// end::verb[]
}
