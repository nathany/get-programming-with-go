package main

import "fmt"

func main() {
	launch := true

	var oneZero int
	if launch {
		oneZero = 1
	} else {
		oneZero = 0
	}
	fmt.Println("Ready for launch:", oneZero)
}
