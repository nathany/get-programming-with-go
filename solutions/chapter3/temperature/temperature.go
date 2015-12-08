package main

import "fmt"

type Celsius float64
type Kelvin float64
type Fahrenheit float64

// KelvinToCelsius  converts ºK to ºC
func KelvinToCelsius(k Kelvin) Celsius { //<1>
	return Celsius(k - 273.15)
}

// Celsius converts ºK to ºC
func (k Kelvin) Celsius() Celsius {
	return Celsius(k - 273.15)
}

// Celsius converts ºF to ºC
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k Kelvin = 294.0
	var c Celsius

	c = KelvinToCelsius(k) //<1>
	c = k.Celsius()        //<2>

	fmt.Print(k, "ºK is ", c, "ºC")
}
