package main

import "fmt"

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k float64) float64 { //<1>
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := KelvinToCelsius(kelvin)         //<2>
	fmt.Print(kelvin, "ºK is ", celsius, "ºC") //<3>
}
