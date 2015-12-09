package main

import "fmt"

func main() {
	a := "text"
	fmt.Printf("%v %T\n", a, a)

	b := 42
	fmt.Printf("%v %T\n", b, b)

	c := 3.14
	fmt.Printf("%v %T\n", c, c)

	d := true
	fmt.Printf("%v %T\n", d, d)
}
