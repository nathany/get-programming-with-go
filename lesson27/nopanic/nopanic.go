package main

import "fmt"

func main() {
	var nowhere *int

	if nowhere != nil {
		fmt.Println(*nowhere)
	}
}
