package main

import "fmt"

func main() {
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
	for _, c := range question {
		fmt.Printf("%c ", c)
	}
}
