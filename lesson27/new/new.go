package main

import "fmt"

func main() {
	var value int
	a := &value
	fmt.Println(*a)

	b := new(int)
	fmt.Println(*b)
}
