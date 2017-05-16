package main

import "fmt"

// Celsius temperatures
type Celsius float64

// Fahrenheit temperatures
type Fahrenheit float64

// Kelvin temperatures
type Kelvin float64

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CelsiusToFahrenheit converts ºC to ºF
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// KelvinToFahrenheit converts ºK to ºF
func KelvinToFahrenheit(k Kelvin) Fahrenheit {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}

func main() {
	var k Kelvin = 294.0
	c := KelvinToCelsius(k)
	fmt.Print(k, "ºK is ", c, "ºC")
}
