package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Kelvin temperature
type Kelvin float64

func measureTemperature(samples int, sensor func() Kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vºK\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() Kelvin {
	return Kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor)
}
