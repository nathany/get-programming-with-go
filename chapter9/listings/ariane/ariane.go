package main

import "fmt"

func main() {
	var bh float64 = 32767
	var h int16 = int16(bh)
	fmt.Println(h)
}
