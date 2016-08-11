package main

import "fmt"

func main() {
	fmt.Println("The year is 2100, should you leap?")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}
