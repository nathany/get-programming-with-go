package main

import "fmt"

func main() {
	const hoursPerDay = 24

	var speed = 100800      // km/h
	var distance = 96300000 // km

	fmt.Println(distance/speed/hoursPerDay, "days")
}
