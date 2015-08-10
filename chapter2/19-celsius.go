package main

import "fmt"

func main() {
	type Celsius float64 //<1>

	var temperature Celsius = 20

	fmt.Println(temperature)
}

// #A The underlying type is float64
