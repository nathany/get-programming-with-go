package main

import "fmt"

func main() {
	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)

	smile := '😃'
	fmt.Printf("%c %[1]v\n", smile)

	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)
}
