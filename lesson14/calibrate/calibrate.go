package main

import "fmt"

// Kelvin temperature
type Kelvin float64

// Sensor function type
type Sensor func() Kelvin

func realSensor() Kelvin {
	return 0
}

func calibrate(sensor Sensor, offset Kelvin) Sensor {
	return func() Kelvin {
		return sensor() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
