package main

import (
	"fmt"
	"math/rand"
)

// Kelvin temperature
type Kelvin float64

func fakeSensor() Kelvin {
	return Kelvin(rand.Intn(151) + 150)
}

func realSensor() Kelvin {
	return 0
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
