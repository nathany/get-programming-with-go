package main

import "fmt"

func main() {
	piggyBank := 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
