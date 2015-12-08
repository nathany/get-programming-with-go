package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" //<1>

func main() {
	year := 2016 //<2>

	switch month := rand.Intn(12) + 1; month { //<3>
	case 2:
		day := rand.Intn(29) + 1 //<4>
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 //<5>
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 //<5>
		fmt.Println(era, year, month, day)
	} //<6>
} //<7>
