package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// celsiusToFahrenheit converts ºC to ºF
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

// kelvinToFahrenheit converts ºK to ºF
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	fmt.Printf("233ºK is %.2fºC\n", kelvinToCelsius(233))
	fmt.Printf("0ºK is %.2fºF\n", kelvinToFahrenheit(0))
}
