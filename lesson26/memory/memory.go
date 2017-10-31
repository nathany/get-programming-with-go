package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)
}
