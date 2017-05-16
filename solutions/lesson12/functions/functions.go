package main

import "fmt"

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// CelsiusToFahrenheit converts ºC to ºF
func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

// KelvinToFahrenheit converts ºK to ºF
func KelvinToFahrenheit(k float64) float64 {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}

func main() {
	fmt.Printf("233ºK is %.2fºC\n", KelvinToCelsius(233))
	fmt.Printf("0ºK is %.2fºF\n", KelvinToFahrenheit(0))
}
