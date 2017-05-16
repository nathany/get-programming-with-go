package main

import "fmt"

// Celsius temperature
type Celsius float64

// Kelvin temperature
type Kelvin float64

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func main() {
	var k Kelvin = 294.0
	c := KelvinToCelsius(k)
	fmt.Print(k, "ºK is ", c, "ºC")
}
