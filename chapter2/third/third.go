package main

import "fmt"

func main() {
	// tag::main[]
	third := 1.0 / 3
	fmt.Println(third)           //<1>
	fmt.Printf("%v\n", third)    //<1>
	fmt.Printf("%f\n", third)    //<2>
	fmt.Printf("%.3f\n", third)  //<3>
	fmt.Printf("%4.2f\n", third) //<4>
	// end::main[]
}

// <1> Both print 0.3333333333333333
// <2> Print 0.333333
// <3> Print 0.333
// <4> Print 0.33
