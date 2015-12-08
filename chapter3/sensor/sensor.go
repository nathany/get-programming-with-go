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
	return 0 //<1>
}

func main() {
	sensor := fakeSensor //<2>
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
