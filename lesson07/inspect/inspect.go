package main

import "fmt"

func main() {
	year := 2016
	fmt.Printf("Type %T for %v.\n", year, year)
	days := 365.2425
	fmt.Printf("Type %T for %[1]v.\n", days)
}
