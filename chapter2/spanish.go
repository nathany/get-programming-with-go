package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")                    //<1>
	fmt.Println(utf8.RuneCountInString(question), "runes") //<2>

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) //<3>
}
