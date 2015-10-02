package main

import "fmt"

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println("T minus", count)
	}
	fmt.Println("Liftoff!")

	for count := 1; count <= 10; count++ {
		fmt.Println("T plus", count)
	}
}
