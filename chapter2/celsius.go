package main

import "fmt"

func main() {
	type Celsius float64

	var temperature Celsius = 20

	fmt.Println(temperature)
}
