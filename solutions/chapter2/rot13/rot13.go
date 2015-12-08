package main

import "fmt"

func main() {
	// message := "I came, I saw, I conquered." // +3

	// message := "abcdefghijklmnopqrstuvwxyzABCDEEFGHIJKLMNOPQRSTUWXYZ"
	// message := "hi international space station"
	message := "Hola Estación Espacial Internacional"
	// message := "привет Международная космическая станция"

	// message := "uv vagreangvbany fcnpr fgngvba"

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

// if c >= 'A' && c <= 'Z' {
//  c = c - 'A'
//  c = (c + 13) % 26
//  c = c + 'A'
// }
