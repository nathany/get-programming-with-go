package main

import "fmt"

func main() {
	launch := false

	var countdown string

	if launch {
		countdown = "Launch in T minus 10 seconds."
	}
	fmt.Println(countdown)
}
