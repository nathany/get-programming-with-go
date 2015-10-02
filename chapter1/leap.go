package main

import "fmt"

func main() {
	year := 2016
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Leap for joy!")
	}
}
