package main

import "fmt"

// celsius temperatures
type celsius float64

// fahrenheit converts ºC to ºF
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// kelvin converts ºC to ºK
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// fahrenheit temperatures
type fahrenheit float64

// celsius converts ºF to ºC
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// kelvin converts ºF to ºK
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

// kelvin temperatures
type kelvin float64

// celsius converts ºK to ºC
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// fahrenheit converts ºK to ºF
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Print(k, "º K is ", c, "º C")
}
