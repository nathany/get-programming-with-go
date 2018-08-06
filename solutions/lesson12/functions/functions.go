package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	fmt.Printf("233º K is %.2fº C\n", kelvinToCelsius(233))
	fmt.Printf("0º K is %.2fº F\n", kelvinToFahrenheit(0))
}
