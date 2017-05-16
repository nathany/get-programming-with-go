package main

import (
	"fmt"
	"math/rand"
)

// Kelvin temperature
type Kelvin float64

// Sensor function type
type Sensor func() Kelvin

func realSensor() Kelvin {
	return 0
}

func fakeSensor() Kelvin {
	return Kelvin(rand.Intn(151) + 150)
}

func calibrate(sensor Sensor, offset Kelvin) Sensor {
	return func() Kelvin {
		return sensor() + offset
	}
}

func main() {
	var offset Kelvin = 5
	sensor := calibrate(fakeSensor, offset)

	for count := 0; count < 10; count++ {
		fmt.Println(sensor())
	}
}
