package main

import "fmt"

func main() {
	const secondsPerDay = 86400

	var speed = 16          // km/s
	var distance = 57600000 // km

	fmt.Println(distance/speed/secondsPerDay, "days")
}
