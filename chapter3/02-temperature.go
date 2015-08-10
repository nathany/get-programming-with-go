package main

import "fmt"

type Celsius float64
type Kelvin float64

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k - 273.15) //#A
}

func main() {
	var k Kelvin = 294.0 //#B
	c := KelvinToCelsius(k)
	fmt.Print(k, "ºK is ", c, "ºC") //#C
}

//#C 294ºK is 20.850000000000023ºC
