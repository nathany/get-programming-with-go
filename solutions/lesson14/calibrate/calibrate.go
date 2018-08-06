package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(fakeSensor, offset)

	for count := 0; count < 10; count++ {
		fmt.Println(sensor())
	}
}
