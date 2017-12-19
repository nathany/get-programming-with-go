package main

import "fmt"

func main() {
	var fn func(a, b int) int
	fmt.Println(fn == nil)

	// fn(1, 2)
}
