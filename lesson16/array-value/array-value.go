package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsMarkII := planets

	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
