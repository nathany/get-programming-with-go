package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ { //#A
		c := message[i]
		if c >= 'a' && c <= 'z' { //#B
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

// #A Iterate through each character
// #B Leave spaces and punctuation as is
