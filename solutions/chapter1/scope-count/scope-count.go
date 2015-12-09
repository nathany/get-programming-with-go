package main

import "fmt"

func main() {
	var count = 0

	for count = 10; count > 0; count-- {
		prefix := "T minus"
		fmt.Println(prefix, count)
	}
	fmt.Println("Liftoff!", count)

	for count := 1; count <= 10; count++ {
		fmt.Println("T plus", count)
	}
}
