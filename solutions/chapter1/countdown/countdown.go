package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println("T minus", count)
		time.Sleep(time.Second)
		count = count - 1
	}
	fmt.Println("Liftoff")
}
