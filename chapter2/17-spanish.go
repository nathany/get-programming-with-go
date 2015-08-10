package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")                    //#A
	fmt.Println(utf8.RuneCountInString(question), "runes") //#B

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) //#C
}

//#A 15 bytes
//#B 12 runes
//#C First rune: ¿ 2
