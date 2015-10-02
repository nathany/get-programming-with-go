package main

import (
	"fmt"
	"strconv"
)

func main() {
	score := 128533

	str := "Score " + strconv.Itoa(score)
	fmt.Println(str)
}
