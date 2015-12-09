package main

import "fmt"

func main() {
	const secondsPerDay = 86400

	var days = 28
	var distance = 56000000 // km

	fmt.Println(distance/(days*secondsPerDay), "km/s")
}
