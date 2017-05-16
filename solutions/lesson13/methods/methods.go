package main

import "fmt"

// Celsius temperatures
type Celsius float64

// Fahrenheit converts ºC to ºF
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// Kelvin converts ºC to ºK
func (c Celsius) Kelvin() Kelvin {
	return Kelvin(c + 273.15)
}

// Fahrenheit temperatures
type Fahrenheit float64

// Celsius converts ºF to ºC
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

// Kelvin converts ºF to ºK
func (f Fahrenheit) Kelvin() Kelvin {
	return f.Celsius().Kelvin()
}

// Kelvin temperatures
type Kelvin float64

// Celsius converts ºK to ºC
func (k Kelvin) Celsius() Celsius {
	return Celsius(k - 273.15)
}

// Fahrenheit converts ºK to ºF
func (k Kelvin) Fahrenheit() Fahrenheit {
	return k.Celsius().Fahrenheit()
}

func main() {
	var k Kelvin = 294.0
	c := k.Celsius()
	fmt.Print(k, "ºK is ", c, "ºC")
}
