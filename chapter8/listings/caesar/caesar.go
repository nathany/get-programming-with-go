package main

import "fmt"

func main() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c)

	c = 'x'
	c = c + 3
	if c > 'z' {
		c = c - 26
	}
	fmt.Printf("%c", c)
}
