package main

import "fmt"

func main() {
	// tag::main[]
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ { //<1>
		c := message[i]
		if c >= 'a' && c <= 'z' { //<2>
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	// end::main[]
}
