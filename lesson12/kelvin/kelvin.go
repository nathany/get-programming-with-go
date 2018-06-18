package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "º K is ", celsius, "º C")
}
