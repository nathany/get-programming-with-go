package main

import (
	"fmt"
	"math/rand"
)

type Kelvin float64

func fakeSensor() Kelvin {
	return Kelvin(rand.Intn(151) + 150)
}

func realSensor() Kelvin {
	return 0 //#A
}

func main() {
	sensor := fakeSensor //#B
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
