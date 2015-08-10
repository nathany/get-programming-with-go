package main

import "fmt"

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third) //#A

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank) //#B
}

// #A 1
// #B 0.30000000000000004
