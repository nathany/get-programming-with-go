package main

import "fmt"

func main() {
	type Celsius float64
	type Fahrenheit float64

	var c Celsius = 20
	var f Fahrenheit = 20

	// invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	// if c == f {
	// }

	// invalid operation: c += f (mismatched types Celsius and Fahrenheit)
	// c += f

	fmt.Println(c, f)
}
