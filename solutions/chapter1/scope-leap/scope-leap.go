package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2016 + rand.Intn(10)
	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
	month := rand.Intn(12) + 1

	days := 31
	switch month {
	case 2:
		days = 28
		if leap {
			days = 29
		}
	case 4, 6, 9, 11:
		days = 30
	}

	day := rand.Intn(days) + 1
	fmt.Println(era, year, month, day)
}
