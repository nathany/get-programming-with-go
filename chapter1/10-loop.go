package main

import "fmt"

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		count = count - 1
	}
}
