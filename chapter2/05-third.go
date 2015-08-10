package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)         //#A
	fmt.Printf("%v ", third)   //#A
	fmt.Printf("%f ", third)   //#B
	fmt.Printf("%.3f ", third) //#C
	fmt.Printf("%4.2f", third) //#D
}

// #A Both print 0.3333333333333333
// #B Print 0.333333
// #C Print 0.333
// #D Print 0.33
