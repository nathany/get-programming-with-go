package main

import "fmt"

func main() {
	var count = 0

	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println(count)
}
