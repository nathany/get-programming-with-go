package main

import "fmt"

func main() {
	var happy rune = 128515
	var sad rune = 128546
	var bang byte = 33

	fmt.Println(happy, sad, bang) //#A

	fmt.Println(string(happy), string(sad), string(bang)) //#A

	fmt.Printf("%c %c %c\n", happy, sad, bang) //#A
}

// #A 128515 128546 33
// #A 😃 😢 !
