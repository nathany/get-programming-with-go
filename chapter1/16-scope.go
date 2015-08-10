package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" //#A

func main() {
	year := 2016 //#B

	switch month := rand.Intn(12) + 1; month { //#C
	case 2:
		day := rand.Intn(29) + 1 //#D
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 //#E
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 //#E
		fmt.Println(era, year, month, day)
	} //#F
} //G

// #A era is available throughout the package
// #B era and year are in scope
// #C era, year, and month are in scope
// #D era, year, month and day are in scope
// #E it’s a new day
// #F month and day are out of scope
// #G year is no longer in scope
