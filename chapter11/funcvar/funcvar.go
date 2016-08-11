package main

import "fmt"

func main() {
	f := func(message string) {
		fmt.Println(message)
	}
	f("Go to the party.")
}
