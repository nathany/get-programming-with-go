package main

import "fmt"

func main() {
	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)

	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)

	d := true
	fmt.Printf("Type %T for %[1]v\n", d)
}
