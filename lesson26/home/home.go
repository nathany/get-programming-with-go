package main

import "fmt"

func main() {
	canada := "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)
}
