package main

import "fmt"

func radiusMap(planet string) float64 {
	r := map[string]float64{
		"Earth": 6371.0,
		"Mars":  3389.5,
	}
	return r[planet]
}

func radiusSwitch(planet string) float64 {
	switch planet {
	case "Earth", "earth":
		return 6371.0
	case "Mars", "mars":
		return 3389.5
	default:
		return 0.0
	}
}

func main() {
	fmt.Println(radiusMap("Earth"))
	fmt.Println(radiusSwitch("Earth"))
}
