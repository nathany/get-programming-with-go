package main

import "fmt"

func main() {
	var happy rune = 128515
	var sad rune = 128546
	var bang byte = 33

	fmt.Println(happy, sad, bang)

	fmt.Println(string(happy), string(sad), string(bang))

	fmt.Printf("%c %c %c\n", happy, sad, bang)
}
