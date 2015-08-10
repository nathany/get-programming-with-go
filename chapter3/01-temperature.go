package main

import "fmt"

// KelvinToCelsius converts ºK to ºC
func KelvinToCelsius(k float64) float64 { //#A
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := KelvinToCelsius(kelvin)         //#B
	fmt.Print(kelvin, "ºK is ", celsius, "ºC") //#C
}

//#B 294ºK is 20.850000000000023ºC
